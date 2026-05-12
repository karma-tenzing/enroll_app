window.onload = function () {
    fetch("/course/all")
        .then(response => response.text())
        .then(data => showAllCourses(data))
}

// create new row in table
function newRow(table, course) {
    var row = table.insertRow(table.length)

    var td = [];
    for (i = 0; i < table.rows[0].cells.length; i++) {
        td[i] = row.insertCell(i);
    }

    // insert data into table cells
    td[0].innerHTML = course.cid
    td[1].innerHTML = course.cname
    td[2].innerHTML = '<input type="button" onclick="deleteCourse(this)" value="delete" id="button-1">'
    td[3].innerHTML = '<input type="button" onclick="updateCourse(this)" value="edit" id="button-2">'
}

// display all courses
function showAllCourses(data) {
    const allCourses = JSON.parse(data)
    var table = document.getElementById("myTable")

    allCourses.forEach(course => {
        newRow(table, course)
    });
}

// display one course
function showCourse(data) {
    const course = JSON.parse(data)
    var table = document.getElementById("myTable")

    newRow(table, course)
}

// reset form
function resetForm() {
    document.getElementById("cid").value = ""
    document.getElementById("cname").value = ""
}

// helper function to get form data
function getFormData() {
    var formData = {
        cid: document.getElementById("cid").value,
        cname: document.getElementById("cname").value
    }

    return formData
}

// add course
function addCourse() {

    data = getFormData()

    // validation
    if (data.cid == "") {
        alert("Course ID can't be empty")
        return
    } else if (data.cname == "") {
        alert("Course Name can't be empty")
        return
    }

    // POST request
    fetch('/course/add', {
        method: "POST",
        body: JSON.stringify(data),
        headers: { "Content-type": "application/json; charset=UTF-8" }
    })
        .then(response1 => {
            if (response1.ok) {

                fetch("/course/" + data.cid)
                    .then(response2 => response2.text())
                    .then(data => showCourse(data))

            } else {
                throw new Error(response1.statusText)
            }
        })
        .catch(e => {
            if (e.message == 401) {
                alert("User not logged in")
                window.open("index.html", "_self")
            } else if (e.message == 400) {
                alert("Bad Request")
            } else {
                alert("Internal server error")
            }
        });

    resetForm()
}

// update API request
function updateAPIRequest(oldCID) {

    newData = getFormData()

    fetch("/course/" + oldCID, {
        method: "PUT",
        body: JSON.stringify(newData),
        headers: { "Content-type": "application/json; charset=UTF-8" }
    })
        .then(res => {
            if (res.ok) {

                selectedRow.cells[0].innerHTML = newData.cid
                selectedRow.cells[1].innerHTML = newData.cname

                // reset button
                var btn = document.getElementById("button-add")
                btn.innerHTML = "Add"
                btn.setAttribute("onclick", "addCourse()")

                selectedRow = null
                resetForm()

            } else {
                alert("server: update request error")
            }
        })
}

var selectedRow = null

// edit course
function updateCourse(input) {

    selectedRow = input.parentElement.parentElement

    document.getElementById("cid").value = selectedRow.cells[0].innerHTML
    document.getElementById("cname").value = selectedRow.cells[1].innerHTML

    cid = selectedRow.cells[0].innerHTML

    // change button to update
    var btn = document.getElementById("button-add")
    btn.innerHTML = "Update"
    btn.setAttribute("onclick", "updateAPIRequest('" + cid + "')")
}

// delete course
function deleteCourse(input) {

    if (confirm("Are you sure to delete this?")) {

        selectedRow = input.parentElement.parentElement
        cid = selectedRow.cells[0].innerHTML

        fetch("/course/" + cid, {
            method: "DELETE"
        })
            .then(res => {
                if (res.ok) {

                    var rowIndex = selectedRow.rowIndex
                    document.getElementById("myTable").deleteRow(rowIndex)

                    alert("Course deleted successfully")

                    selectedRow = null

                } else {
                    alert("server: delete request error")
                }
            })
    }
}