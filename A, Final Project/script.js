function redirectToModule(module) {
    switch (module) {
        case 'patient':
            alert("Redirecting to Patient Module...");
            
            break;
        case 'doctor':
            alert("Redirecting to Doctor Module...");
            
            break;
        case 'appointment':
            alert("Redirecting to Appointment Module...");
            
            break;
        case 'diagnostics':
            alert("Redirecting to Diagnostics Module...");
            
            break;
        case 'billing':
            alert("Redirecting to Billing & Payment Module...");
            
            break;
        default:
            alert("Invalid module selected.");
    }
}