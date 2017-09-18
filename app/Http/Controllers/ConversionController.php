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

        return view('converters', [
            'word' => $word,
            'number' => $number
        ]);
    }

    public function word_to_num ($input = null)
    {
        if ($input === null) return response()->json(['result' => []]);

        $words = $this->split($input);

        foreach ($words as $word) {
            $num = Word::distinct()
                ->select('number')
                ->where('word', $word)
                ->orderBy('number', 'asc')
                ->get()
                ->pluck('number')
                ->toArray();

            if (array_filter($num, 'strlen')) {
                if (count($num) == 1) {
                    $num = array_shift($num);
                }
            } else {
                if (!isset($letters)) {
                    $resp = Letter::select('number', 'symbol')
                        ->orderBy(DB::raw('length(`symbol`)'), 'desc')
                        ->get();

                    $letters = [
                        'number' => $resp->pluck('number')->toArray(),
                        'symbol' => $resp->pluck('symbol')->toArray()
                    ];
                }

                // Replace every symbol with its corresponding number
                $num = str_replace($letters['symbol'], $letters['number'], strtolower($word));

                // In the end, remove all non-digits with regex
                $num = preg_replace('/[^\d]/', '', $num);
            }

            $nums[$word] = $num;
        }

        return response()->json(['result' => $nums]);
    }

    public function num_to_word ($input = null)
    {
        if ($input === null) return response()->json(['result' => []]);

        $nums = $this->split($input);

        foreach ($nums as $num) {
            $word = Word::where('number', $num)
                ->distinct()
                ->select('word')
                ->orderBy('word', 'asc')
                ->get()
                ->pluck('word')
                ->toArray();

            if (count($word) == 1) {
                $word = array_shift($word);
            }

            $words[$num] = $word;
        }

        return response()->json(['result' => $words]);
    }

    private function split($input) {
        return preg_split('/[,;\s]/', $input, NULL, PREG_SPLIT_NO_EMPTY);
    }
}
