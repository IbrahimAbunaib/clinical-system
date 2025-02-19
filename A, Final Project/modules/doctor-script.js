function showSection(sectionId) {

    document.querySelectorAll('.section').forEach(section => {
        section.classList.add('hidden');
    });

    document.getElementById(sectionId).classList.remove('hidden');
}

function saveData(event, type) {
    event.preventDefault(); 

    const formData = new FormData(event.target); 
    fetch('save_data.php', {
        method: 'POST',
        body: formData
    })
    .then(response => response.text())
    .then(data => {
        alert(data); 
        displaySavedData(type, formData); 
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

function fetchData(type, sectionId) {
    fetch(`get_data.php?type=${type}`)
        .then(response => response.text())
        .then(data => {
            document.getElementById(sectionId).innerHTML = data;
        });
}

function displaySavedData(type, formData) {
    const savedDataDiv = document.getElementById(`saved${type.charAt(0).toUpperCase() + type.slice(1)}`);
    savedDataDiv.innerHTML = `<h3>Saved Data:</h3>`;

    for (let [key, value] of formData.entries()) {
        savedDataDiv.innerHTML += `<p><strong>${key}:</strong> ${value}</p>`;
    }
}