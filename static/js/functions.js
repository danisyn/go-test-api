function changePath() {
    print("Activated")
    var actual_path = window.location.pathname;
    return actual_path.replace(actual_path, "/api/logs")
}