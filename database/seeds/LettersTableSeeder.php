<?php

use Illuminate\Database\Seeder;

use App\Letter;

class LettersTableSeeder extends Seeder
{

    /**
     * Auto generated seed file
     *
     * @return void
     */
    public function run()
    {
        Letter::truncate();

        Letter::insert([
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
                'symbol' => 'th',
            ],
            [
                'number' => 1,
                'symbol' => 't',
            ],
            [
                'number' => 1,
                'symbol' => 'd',
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
                'symbol' => 'ch',
            ],
            [
                'number' => 6,
                'symbol' => 'j',
            ],
            [
                'number' => 6,
                'symbol' => 'sh',
            ],
            [
                'number' => 7,
                'symbol' => 'ck',
            ],
            [
                'number' => 7,
                'symbol' => 'c',
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
                'number' => 7,
                'symbol' => 'q',
            ],
            [
                'number' => 8,
                'symbol' => 'f',
            ],
            [
                'number' => 8,
                'symbol' => 'ph',
            ],
            [
                'number' => 8,
                'symbol' => 'v',
            ],
            [
                'number' => 9,
                'symbol' => 'b',
            ],
            [
                'number' => 9,
                'symbol' => 'p',
            ],
        ]);
    }
}
