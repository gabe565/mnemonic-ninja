<?php

use Illuminate\Database\Seeder;

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

        $db_resp = DB::table('ipas')
            ->select('number', 'symbol')
            ->orderBy(DB::raw('length(`symbol`)'), 'desc')
            ->get()
            ->toArray();
        $ipas['symbol'] = array_column($db_resp, 'symbol');
        $ipas['number'] = array_column($db_resp, 'number');

        DB::table('words')->truncate();
        DB::connection()->disableQueryLog();

        $handle = fopen(base_path('resources/cmudict/cmudict.dict'), 'r');

        $result = [];
        $n = 0;
        while (($line = fgets($handle)) !== false) {
            $entry = explode(' ', trim($line), 2);

            // Convert to lowercase, remove (#) on duplicates into `word`
            $word = strtok($entry[0], '(');//preg_replace('/\(\d\)/', '', $entry[0]);
            // Remove comments and convert to IPA into `ipa`
            $ipa = $arpabetToIPA->getIPA(strtok($entry[1], '#'));

            // Number calculated from the list of symbols which are fetched above.
            $number = filter_var(str_replace($ipas['symbol'], $ipas['number'], $ipa), FILTER_SANITIZE_NUMBER_INT);

            $result[] = [
                'word' => $word,
                'ipa' => $ipa,
                'number' => $number,
            ];

            ++$n;
            if ($n > 500) {
                DB::table('words')->insert($result);
                $result = [];
                $n = 0;
            }
        }

        DB::table('words')->insert($result);

        // Set any empty numbers to null
        DB::table('words')
            ->where('number', '')
            ->update(['number' => null]);
    }
}
