<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateIpasTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('ipas', function (Blueprint $table) {
            $table->smallIncrements('id');
            $table->tinyInteger('number')->unsigned();
            $table->string('symbol', 4);
            $table->tinyInteger('length')->nullable()->virtualAs('length(`symbol`)');
            $table->index('length', 'length');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('ipas');
    }
}
