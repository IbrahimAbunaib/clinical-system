function showSection(sectionId) {
    
    document.querySelectorAll('.section').forEach(section => {
        section.classList.add('hidden');
    });

    document.getElementById(sectionId).classList.remove('hidden');
}

function bookAppointment() {
    alert("Booking appointment...");
}

function rescheduleAppointment() {
    alert("Rescheduling appointment...");
}

function cancelAppointment() {
    alert("Canceling appointment...");
}

function viewHistory() {
    alert("Viewing medical history...");
}

function viewNotifications() {
    alert("Viewing notifications...");
}