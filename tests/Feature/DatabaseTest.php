<?php

namespace Tests\Feature;

use Tests\TestCase;
use Illuminate\Foundation\Testing\RefreshDatabase;

class DatabaseTest extends TestCase
{
    /**
     * A test to make sure the database is not empty.
     *
     * @return void
     */
    public function testToNum()
    {
        $this->assertDatabaseHas('words', [
            'word' => 'example',
            'number' => '70395',
        ]);
    }
}
