<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Word;
use App\Letter;

class ConversionController extends Controller
{
    public function word_to_num ($word = null)
    {
        if ($word === null) return [];

        $nums = Word::where('word', $word)
            ->distinct()
            ->select('number')
            ->orderBy('number', 'asc')
            ->get()
            ->pluck('number')
            ->toArray();

        if (array_filter($nums)) {
            return response()->json($nums);
        }
        else {
            $resp = Letter::select('number', 'symbol')
                ->orderBy('length', 'desc')
                ->get();

            $letters = [
                'number' => $resp->pluck('number')->toArray(),
                'symbol' => $resp->pluck('symbol')->toArray()
            ];

            // Replace every symbol with its corresponding number
            $nums = str_replace($letters['symbol'], $letters['number'], $word);

            // In the end, remove all non-digits with regex
            $nums = preg_replace('/[^\d]/', '', $nums);

            return response()->json([$nums]);
        }
    }

    public function num_to_word ($num = null)
    {
        if ($num === null) return [];

        $words = Word::where('number', $num)
            ->distinct()
            ->select('word')
            ->orderBy('word', 'asc')
            ->get()
            ->pluck('word')
            ->toArray();

        return response()->json($words);
    }
}
