<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DB;
use App\Word;
use App\Letter;

class ConversionController extends Controller
{

    public function show ()
    {
        $word = $_GET['word'] ?? '';
        $number = $_GET['number'] ?? '';

        return view('body', [
            'word' => $word,
            'number' => $number
        ]);
    }

    public function word_to_num ($input = null)
    {
        if ($input === null)
            return response()->json(['result' => []]);

        // Split input on any whitespace, comma, or semicolon
        $words = $this->split($input);

        foreach ($words as $word) {
            // Get all numbers for the current word
            $num = Word::distinct()
                ->select('number')
                ->where('word', $word)
                ->orderBy('number', 'asc')
                ->get()
                ->pluck('number')
                ->toArray();

            // If $num is empty then we will guess by letter
            if (empty($num)) {
                if (!isset($letters)) {
                    // Get letters and numbers from the database
                    $resp = Letter::select('number', 'symbol')
                        ->orderBy(DB::raw('length(`symbol`)'), 'desc')
                        ->get();

                    // Rearrange array into two arrays which can be passed to str_replace
                    $letters = [
                        'symbol' => $resp->pluck('symbol')->toArray(),
                        'number' => $resp->pluck('number')->toArray()
                    ];
                }

                // Remove minus symbol since filter_var will not pick it up
                $letters['symbol'][] = '-';
                $letters['number'][] = '';

                // Replace every symbol with its corresponding number
                $guess = str_replace($letters['symbol'], $letters['number'], strtolower($word));

                // Remove all non-digits and add to $num
                $num[] = filter_var($guess, FILTER_SANITIZE_NUMBER_INT);
            }

            $nums[] = [
                'q' => $word,
                'r' => $num
            ];
        }

        return response()->json([
            'result' => $nums
        ]);
    }

    public function num_to_word ($input = null)
    {
        if ($input === null)
            return response()->json(['result' => []]);

        // Split input on any whitespace, comma, or semicolon
        $nums = $this->split($input);

        foreach ($nums as $num) {
            // Get all words for the current num
            $word = Word::where('number', $num)
                ->distinct()
                ->select('word')
                ->orderBy('word', 'asc')
                ->get()
                ->pluck('word')
                ->toArray();

            $words[] = [
                'q' => $num,
                'r' => $word
            ];
        }

        return response()->json([
            'result' => $words
        ]);
    }

    private function split($input) {
        return preg_split('/[,;\s]/', $input, NULL, PREG_SPLIT_NO_EMPTY);
    }
}
