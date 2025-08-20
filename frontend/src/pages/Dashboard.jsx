import { useEffect, useState } from "react";
import api from "../api";
import TaskList from "../components/TaskList";
import TaskForm from "../components/TaskForm";

function Dashboard() {
  const [tasks, setTasks] = useState([]);
  const [showForm, setShowForm] = useState(false);
  const token = localStorage.getItem("token")
  const handleLogout = () => {
    localStorage.removeItem("token");
    window.location.href = "/login";
  };
  
  const fetchTasks = async () => {
    console.log("Token:", token)
    const res = await api.get("http://localhost:8080/api/tasks");
    setTasks(res.data);
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const closeModal = () => setShowForm(false);

  return (
    <div className="min-h-screen bg-gray-100 py-8">
      <div className="max-w-2xl mx-auto px-4">
        {/* Header */}
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-bold text-gray-800">My Tasks</h1>
          <button
            onClick={handleLogout}
            className="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600 transition duration-200"
          >
            Log Out
          </button>
        </div>

        {/* Add Task Button */}
        

        {/* Task List */}
        <TaskList tasks={tasks} refresh={fetchTasks} />
      </div>
      <br />
      <div className="flex padding-2 justify-center mb-4">
          <button
            onClick={() => setShowForm(true)}
            className="bg-black text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition duration-200"
          >
            + Add Task
          </button>
        </div>
      {/* Modal */}
      {showForm && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg shadow-lg p-6 w-full max-w-lg relative">
            {/* Close Button */}
            <button
              onClick={closeModal}
              className="absolute top-3 right-3 text-gray-500 hover:text-gray-700"
            >
              ✖
            </button>

            {/* Task Form */}
            <TaskForm
              refresh={() => {
                fetchTasks();
                closeModal();
              }}
            />
          </div>
        </div>
      )}
    </div>
  );
}

export default Dashboard;
