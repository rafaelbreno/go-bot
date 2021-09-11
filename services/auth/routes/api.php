<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::post('/create', [\App\Http\Controllers\UserController::class, 'store']);
Route::post('/login', [\App\Http\Controllers\UserController::class, 'login']);
Route::get('/profile', [\App\Http\Controllers\UserController::class, 'index'])
    ->middleware(['auth:sanctum']);


Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});
