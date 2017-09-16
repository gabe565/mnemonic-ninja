<?php

namespace Tests\Feature;

use Tests\TestCase;
use Illuminate\Foundation\Testing\RefreshDatabase;

class JSONTest extends TestCase
{
    /**
     * Test word to number conversion
     *
     * @return void
     */
    private function toNum($word, $num)
    {
        if (!is_array($num)) {
            $num = [$num];
        }
        $response = $this->json('GET', "/api/to/num/$word");
        $response
            ->assertStatus(200)
            ->assertJson([
                'result' => $num,
            ]);
    }

    /**
     * Test number to word conversion
     *
     * @return void
     */
    private function toWord($num, $word)
    {
        if (!is_array($word)) {
            $word = [$word];
        }
        $response = $this->json('GET', "/api/to/word/$num");
        $response
            ->assertStatus(200)
            ->assertJson([
                'result' => $word,
            ]);
    }

    public function testAPI() {
        $this->toNum('example', 70395);
        $this->toWord(70395, 'example');

        $this->toNum('garage', 746);
        $this->toNum('garages', 7460);
        $this->toNum('garagez', 7470);

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
        $this->toNum('.h .y .w a e i o u', []);
    }
}
