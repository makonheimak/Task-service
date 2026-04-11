<?php

declare(strict_types=1);

namespace App\Domain\Task\DTO;

final class TaskData
{
    public function __construct(
        public readonly ?int $id,
        public readonly string $task,
        public readonly int $userId,
        public readonly bool $isDone = false,
    ) {
    }

    /**
     * @param array<string, mixed> $payload
     */
    public static function fromArray(array $payload): self
    {
        return new self(
            id: isset($payload['id']) ? (int) $payload['id'] : null,
            task: (string) $payload['task'],
            userId: (int) $payload['user_id'],
            isDone: (bool) ($payload['is_done'] ?? false),
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function toArray(): array
    {
        return [
            'id' => $this->id,
            'task' => $this->task,
            'user_id' => $this->userId,
            'is_done' => $this->isDone,
        ];
    }
}
