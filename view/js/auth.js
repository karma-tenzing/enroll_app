if(document.cookie == ""){
    alert("User is not logged in!!!")
    window.open("index.html", "_self")
}else{
    console.log("cookie set")
}