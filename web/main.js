document.addEventListener("DOMContentLoaded", function () {
  loadTasks();

  const statuses = ["todo", "in-progress", "done"];

  async function loadTasks() {
    const response = await fetch('/tasks');
    const tasks = await response.json();

    tasks.forEach(task => {
      const taskElement = createTaskElement(task);
      document.getElementById(`tasks-${statuses[task.status]}`).appendChild(taskElement);
    });
  }

  function createTaskElement(task) {
    const taskElement = document.createElement('div');
    taskElement.classList.add('task');
    taskElement.textContent = task.title;
    taskElement.draggable = true;
    taskElement.ondragstart = (e) => handleDragStart(e, task);
    return taskElement;
  }

  function handleDragStart(e, task) {
    e.dataTransfer.setData('task', JSON.stringify(task));
  }

  document.querySelectorAll('.tasks').forEach(taskContainer => {
    taskContainer.ondragover = (e) => e.preventDefault();
    taskContainer.ondrop = (e) => handleDrop(e);
  });

  function handleDrop(e) {
    e.preventDefault();
    const task = JSON.parse(e.dataTransfer.getData('task'));
    const newStatus = statuses.indexOf(e.target.id.split('-')[1]);

    task.status = newStatus;
    updateTask(task);
  }

  async function updateTask(task) {
    await fetch(`/tasks/${task.ID}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(task)
    });
    loadTasks();
  }

  window.showForm = (status) => {
    document.getElementById('task-form').style.display = 'block';
    document.getElementById('task-status').value = statuses.indexOf(status);
  };

  window.closeForm = () => {
    document.getElementById('task-form').style.display = 'none';
  };

  window.addTask = async () => {
    const title = document.getElementById('task-title').value;
    const description = document.getElementById('task-description').value;
    const status = parseInt(document.getElementById('task-status').value, 10);

    const newTask = { title, description, status };
    await fetch('/tasks', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(newTask)
    });

    document.getElementById('task-title').value = '';
    document.getElementById('task-description').value = '';

    closeForm();
    loadTasks();
  };
});
