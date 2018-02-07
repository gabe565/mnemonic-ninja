<?php

class DatabaseTest extends TestCase
{
    /**
     * A test to make sure the database is not empty.
     *
     * @return void
     */
    public function testToNum()
    {
        $this->seeInDatabase('words', [
            'word' => 'example',
            'number' => '70395',
        ]);
    }
}
