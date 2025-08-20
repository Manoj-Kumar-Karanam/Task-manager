import { useState } from "react";
import api from "../api";

function TaskForm({ refresh }) {
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");
  const [completed, setCom] = useState(false);

  const handleSubmit = async () => {
    if (!title.trim()) {
      console.error("Cannot create task with empty title");
      return;
    }

    await api.post("/api/tasks", { title, description: desc, completed });
    setTitle("");
    setDesc("");
    setCom(false);
    refresh();
  };

  const handleCompletedChange = (e) => {
    setCom(e.target.value === "Yes");
  };

  return (
    <div className="max-w-md mx-auto bg-white shadow-md rounded-lg p-6 space-y-4">
      <h2 className="text-xl font-semibold text-gray-800 text-center">
        Add New Task
      </h2>

      <input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Title"
        className="w-full px-4 py-2 border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400"
      />

      <textarea
        value={desc}
        onChange={(e) => setDesc(e.target.value)}
        placeholder="Description"
        rows="3"
        className="w-full px-4 py-2 border-2 border-gray-300 rounded-lg resize-none focus:outline-none focus:ring-2 focus:ring-blue-400"
      />

      <div>
        <label
          htmlFor="comp"
          className="block mb-2 text-sm font-medium text-gray-700"
        >
          Completed
        </label>
        <select
          id="comp"
          value={completed ? "Yes" : "No"}
          onChange={handleCompletedChange}
          className="w-full px-4 py-2 border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400"
        >
          <option value="Yes">Yes</option>
          <option value="No">No</option>
        </select>
      </div>

      <button
        id="task-title"
        onClick={handleSubmit}
        className="w-full bg-blue-600 text-white py-2 rounded-lg hover:bg-blue-700 transition duration-200"
      >
        Add Task
      </button>
    </div>
  );
}

export default TaskForm;
