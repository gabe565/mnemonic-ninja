<?php

use Illuminate\Database\Seeder;

class IpasTableSeeder extends Seeder
{

    /**
     * Auto generated seed file
     *
     * @return void
     */
    public function run()
    {
        DB::table('ipas')->truncate();

        DB::table('ipas')->insert([
            [
                'number' => 0,
                'symbol' => 's',
            ],
            [
                'number' => 0,
                'symbol' => 'z',
            ],
            [
                'number' => 1,
                'symbol' => 'ð',
            ],
            [
                'number' => 1,
                'symbol' => 'θ',
            ],
            [
                'number' => 1,
                'symbol' => 'd',
            ],
            [
                'number' => 1,
                'symbol' => 't',
            ],
            [
                'number' => 27,
                'symbol' => 'ŋ',
            ],
            [
                'number' => 2,
                'symbol' => 'n',
            ],
            [
                'number' => 3,
                'symbol' => 'm',
            ],
            [
                'number' => 4,
                'symbol' => 'r',
            ],
            [
                'number' => 5,
                'symbol' => 'l',
            ],
            [
                'number' => 6,
                'symbol' => 'dʒ',
            ],
            [
                'number' => 6,
                'symbol' => 'ʒ',
            ],
            [
                'number' => 6,
                'symbol' => 'tʃ',
            ],
            [
                'number' => 6,
                'symbol' => 'ʃ',
            ],
            [
                'number' => 7,
                'symbol' => 'g',
            ],
            [
                'number' => 7,
                'symbol' => 'k',
            ],
            [
                'number' => 8,
                'symbol' => 'f',
            ],
            [
                'number' => 9,
                'symbol' => 'b',
            ],
            [
                'number' => 8,
                'symbol' => 'v',
            ],
            [
                'number' => 9,
                'symbol' => 'p',
            ],
        ]);
    }
}
