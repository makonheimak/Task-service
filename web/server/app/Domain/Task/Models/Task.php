<?php

declare(strict_types=1);

namespace App\Domain\Task\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

final class Task extends Model
{
    use HasFactory;
    use SoftDeletes;

    protected $table = 'tasks';

    protected $fillable = [
        'task',
        'user_id',
        'is_done',
    ];

    protected $casts = [
        'is_done' => 'bool',
    ];
}
