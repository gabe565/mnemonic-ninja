<?php

class JSONTest extends TestCase
{
    /**
     * Test word to number conversion
     *
     * @return void
     */
    private function toNum($query, $num)
    {
        $words = explode(' ', $query);
        foreach ($words as $key => $word) {
            $response = $this->json('GET', "/api/to/num/$word");
            if (is_array($num)) {
                $response->seeJson([
                    'q' => $word,
                    'r' => ["${num[$key]}"]
                ]);
            } else {
                $response->seeJson([
                    'q' => $query,
                    'r' => ["$num"]
                ]);
            }
        }
    }

    /**
     * Test number to word conversion
     *
     * @return void
     */
    private function toWord($num, $word)
    {
        $response = $this->json('GET', "/api/to/word/$num");
        $response
            ->seeJson([
                'q' => "$num",
                'r' => [$word]
            ]);
    }

    public function testAPI() {
        $this->toNum('example', '70395');
        $this->toWord('70395', 'example');

        // Test a word in wordlist, then without
        $this->toNum('garage', '746');
        $this->toNum('garages', '7460');
        $this->toNum('garagez', '7470');

        // Test letter conversions
        $this->toNum('s z', [0, 0]);
        $this->toNum('t d .th', [1, 1, 1]);
        $this->toNum('n', 2);
        $this->toNum('m', 3);
        $this->toNum('r', 4);
        $this->toNum('l', 5);
        $this->toNum('ch j sh', [6, 6, 6]);
        $this->toNum('ck .c .g k q', [7, 7, 7, 7, 7]);
        $this->toNum('f .ph v', [8, 8, 8]);
        $this->toNum('b p', [9, 9]);
        $this->toNum('.h .y .w a e i o u', ['', '', '', '', '', '', '', '']);

        // Test to hit all of the IPA conversions
        $this->toNum('antidisestablishmentarianism', '2110019563214203');
        $this->toWord('2110019563214203', 'antidisestablishmentarianism');
        $this->toNum('genossenschaftsbank', '720268109277');
        $this->toWord('720268109277', 'genossenschaftsbank');
    }
}
