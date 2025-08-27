// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.0 <0.9.0;

contract Todo {
    struct Task {
        string description;
        bool completed;
    }

    Task[] public tasks;

    constructor() {}

    event TaskAdded(uint indexed taskId, string description);
    event TaskCompleted(uint indexed taskId);

    function addTask(string memory _description) public {
        tasks.push(Task({description: _description, completed: false}));
        emit TaskAdded(tasks.length - 1, _description);
    }

    function completeTask(uint _taskId) public {
        require(_taskId < tasks.length, "Task does not exist");
        require(!tasks[_taskId].completed, "Task already completed");
        tasks[_taskId].completed = true;
        emit TaskCompleted(_taskId);
    }

    function getTask(uint _taskId) public view returns (string memory, bool) {
        require(_taskId < tasks.length, "Task does not exist");
        Task memory task = tasks[_taskId];
        return (task.description, task.completed);
    }

    function getAllTasks() public view returns (Task[] memory) {
        return tasks;
    }
}
