window.onload = function (){
    fetch("/student/all")
    .then(response => response.text())
    .then(data => showAllStudents(data))
}

function newRow(table, student){
    var row = table.insertRow(table.length)
        
        var td =[];
        for(i = 0; i<table.rows[0].cells.length; i++){
            td[i] = row.insertCell(i);
        }

        // insert data in td
        td[0].innerHTML = student.stdid
        td[1].innerHTML = student.fname
        td[2].innerHTML = student.lname
        td[3].innerHTML = student.email
        td[4].innerHTML = '<input type="button" onclick="deleteStudent(this)" value="delete" id="button-1">'
        td[5].innerHTML = '<input type="button" onclick="updateStudent(this)" value="edit" id="button-2">'

}

function showAllStudents(data){
    const allStudents = JSON.parse(data)
    var table = document.getElementById("myTable")
    allStudents.forEach(stud => {
        newRow(table, stud)
    });

}


function showStudent(data){
    // console.log(data);
    // convert json string to js object 
    const student = JSON.parse(data)
    var table = document.getElementById("myTable")
    newRow(table, student)



        
}
// form reset 
function resetForm(){
    document.getElementById("sid").value = ""
    document.getElementById("fname").value = ""
    document.getElementById("lname").value = ""
    document.getElementById("email").value = ""
}
function addStudent(){
    // create js object to store form data
    var data = {
        stdid : parseInt(document.getElementById("sid").value),
        fname : document.getElementById("fname").value,
        lname : document.getElementById("lname").value,
        email : document.getElementById("email").value
    }
    // form validation

    var sid = data.stdid
    if(isNaN(sid)){
        alert("enter valid student ID")
        return
    }else if(data.email == ""){
        alert("Email cant be empty")
        return
    }else if(data.fname == ""){
        alert("First name cant be empty")
        return
    }
        
    // call POST API
    fetch('/student/add', {
        method: "POST",
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response1 => {
        // check response from the fetch 
        if(response1.ok){
            fetch("/student/"+data.stdid)
            .then(response2 => response2.text())
            .then(data => showStudent(data))
        }else{
            throw new Error(response1.statusText)
        }
    }).catch(e => alert(e));
    resetForm()
}

function updateStudent(input){
    // get the selected row
    var selectedRow = input.parentElement.parentElement
    document.getElementById("sid").value = selectedRow.cells[0].innerHTML
    document.getElementById("fname").value = selectedRow.cells[1].innerHTML
    document.getElementById("lname").value = selectedRow.cells[2].innerHTML
    document.getElementById("email").value = selectedRow.cells[3].innerHTML

    // change button value to update
    var btn = document.getElementById("button-add")
    btn.innerHTML = "Update"
    btn.setAttribute("onclick", "updateAPIRequest(sid)")
}