<?php

declare(strict_types=1);

namespace App\Domain\Task\Http\Controllers;

use App\Domain\Task\Http\Requests\StoreTaskRequest;
use App\Domain\Task\Http\Requests\UpdateTaskRequest;
use App\Domain\Task\Models\Task;
use Illuminate\Http\JsonResponse;

final class TaskController
{
    public function index(): JsonResponse
    {
        return response()->json(Task::query()->latest('id')->get());
    }

    public function store(StoreTaskRequest $request): JsonResponse
    {
        $task = Task::query()->create($request->validated());

        return response()->json($task, 201);
    }

    public function show(int $id): JsonResponse
    {
        $task = Task::query()->findOrFail($id);

        return response()->json($task);
    }

    public function update(UpdateTaskRequest $request, int $id): JsonResponse
    {
        $task = Task::query()->findOrFail($id);
        $task->update($request->validated());

        return response()->json($task->refresh());
    }

    public function destroy(int $id): JsonResponse
    {
        Task::query()->findOrFail($id)->delete();

        return response()->json(status: 204);
    }
}
