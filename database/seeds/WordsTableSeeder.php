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
            ->select('number','symbol')
            ->orderBy('length', 'desc')
            ->get()
            ->toArray();
        $ipas['symbol'] = array_column($db_resp, 'symbol');
        $ipas['number'] = array_column($db_resp, 'number');

        DB::table('words')->truncate();
        DB::connection()->disableQueryLog();

        $handle = fopen('https://raw.githubusercontent.com/cmusphinx/cmudict/master/cmudict.dict', 'r');

        $result = [];
        $n = 0;
        while (($line = fgets($handle)) !== false) {
            $entry = explode(' ', $line, 2);

            // Convert to lowercase, remove (#) on duplicates into `word`
            $word = preg_replace('/\(\d+?\)/', '', strtolower($entry[0]));
            // Remove newlines from end, remove comments, and convert to IPA into `ipa`
            $ipa = $arpabetToIPA->getIPA(preg_replace(["/\n/", "/#(.*)/"], '', $entry[1]));

            // Number calculated from the list of symbols which are fetched above.
            $number = preg_replace('/[^\d]/', '',
                str_replace($ipas['symbol'], $ipas['number'], $ipa)
            );

            if ($number === '') {
                $number = null;
            }

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
    }
}
