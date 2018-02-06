<?php

use Illuminate\Database\Seeder;
use App\Ipa;
use App\Word;

class WordsTableSeeder extends Seeder
{

    /**
     * Auto generated seed file
     *
     * @return void
     */
    public function run()
    {
        $arpabetToIPA = new ArpabetToIPA\App();

        $db_resp = Ipa::orderBy(DB::raw('length(`symbol`)'), 'desc')
            ->get();
        $ipas = [
            'number' => $db_resp->pluck('number')->all(),
            'symbol' => $db_resp->pluck('symbol')->all()
        ];

        // Open cmudict
        $handle = fopen(base_path('resources/cmudict/cmudict.dict'), 'r');

        // Parse through cmudict and store into a collection
        $collection = collect([]);
        while (($line = fgets($handle)) !== false) {
            // Separate into word and 
            $word = strtok(trim($line), ' ');
            $arpabet = strtok('');

            // Remove (#) from duplicates
            $word = strtok($word, '(');

            // Remove comments and convert to IPA
            $ipa = $arpabetToIPA->getIPA(strtok($arpabet, '#'));

            // Number calculated from the list of symbols, then extra letters filtered out
            $number = filter_var(str_replace($ipas['symbol'], $ipas['number'], $ipa), FILTER_SANITIZE_NUMBER_INT);

            $collection[] = [
                'word' => $word,
                'ipa' => $ipa,
                'number' => $number,
            ];
        }
        fclose($handle);

        Word::truncate();

        // Insert 500 into the database at a time
        foreach($collection->chunk(500) as $chunk) {
            Word::insert($chunk->all());
        }

        // Set any empty numbers to null
        Word::where('number', '')
            ->update(['number' => null]);
    }
}
