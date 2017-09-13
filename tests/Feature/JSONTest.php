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
    public function testToNum()
    {
        $response = $this->json('GET', '/api/to/num/example');
        $response
            ->assertStatus(200)
            ->assertJson([
                'result' => ['70395'],
            ]);
    }

    /**
     * Test number to word conversion
     *
     * @return void
     */
    public function testToWord()
    {
        $response = $this->json('GET', '/api/to/word/70395');
        $response
            ->assertStatus(200)
            ->assertJson([
                'result' => ['example'],
            ]);
    }

}
