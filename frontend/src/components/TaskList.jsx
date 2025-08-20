import { useState } from "react";
import api from "../api";

function TaskList({ tasks, refresh }) {
  const [selectedTask, setSelectedTask] = useState(null); // store clicked task

  const deleteTask = async (ID) => {
    if (window.confirm("Are you sure you want to delete?")) {
      await api.delete(`/api/tasks/${ID}`);
      refresh();
    }
  };

  return (
    <div className="max-w-md mx-auto bg-white shadow-md rounded-lg p-6">
      <h2 className="text-xl font-semibold text-gray-800 mb-4 text-center">
        Task List
      </h2>

      {tasks.length === 0 ? (
        <p className="text-gray-500 text-center">No tasks available.</p>
      ) : (
        <ul className="space-y-3">
          {tasks.map((task) => (
            <li
              key={task.ID}
              onClick={() => setSelectedTask(task)} // open modal
              className="flex items-center justify-between bg-gray-50 border border-gray-200 rounded-lg px-4 py-2 hover:bg-gray-100 cursor-pointer"
            >
              <span
                className={task.completed ? "line-through text-gray-500" : ""}
              >
                {task.Title}
              </span>
              <button
                onClick={(e) => {
                  e.stopPropagation(); // prevent modal open when clicking delete
                  deleteTask(task.ID);
                }}
                className="text-red-500 hover:text-red-700 transition-colors cursor-pointer"
              >
                Delete
              </button>
            </li>
          ))}
        </ul>
      )}

      {/* Modal */}
      {selectedTask && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-40 z-50">
          <div className="bg-white rounded-xl shadow-lg max-w-sm w-full p-6 relative">
            <button
              className="absolute top-2 right-2 text-gray-500 hover:text-gray-700"
              onClick={() => setSelectedTask(null)}
            >
              ✕
            </button>
            <h3 className="text-lg font-semibold mb-2 text-gray-800">
              {selectedTask.Title}
            </h3>
            <p className="text-gray-600 mb-2">
              <strong>Status:</strong>{" "}
              {selectedTask.completed ? "Completed" : "Pending"}
            </p>
            <p className="text-gray-600 mb-2">
              <strong>ID:</strong> {selectedTask.ID}
            </p>
            {/* If you have more fields, show them here */}
          </div>
        </div>
      )}
    </div>
  );
}

export default TaskList;
