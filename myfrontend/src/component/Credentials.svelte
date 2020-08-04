<script>
import { Table , Alert, Button} from 'sveltestrap';
import axios       from 'axios';
import Cookies     from 'js-cookie';

let cookieName   = "session";
let btncolor     = "light";
let failcolor    = "danger";
let succolor     = "success";
let visible      = true;
let table        = { empty : true , updated : false};
let content      = [];
let CNT          = [];
let fieldsRequire= false;
let copied       = false;
let payload      = {Category:"",Credential:[]};
let saveBtndisbl = true;
var numObj       = 4;
var list         = new Array(2*numObj);
let filtered;

function CheckCookie(){
  var cookie = Cookies.get(cookieName);
  if (cookie != null) {
    return true;
  } else {
    return false ;
  }
}


function addToList(item,index) {
  CNT.push(content[index].Username);
  CNT.push(content[index].Password);
  CNT.push(content[index].Category);
}

function clearTable() {
  var table = document.getElementById("myrows");
  var inp   = table.getElementsByTagName("input");
  for (var i=0;i<inp.length;i++) {
    inp[i].value = '';
  }
}

function GetTableData() {
  var category = document.getElementById("ctg");
  var bodyFormData = new FormData();
  filtered = category.value;
  bodyFormData.set("category",category.value);
  var cookieExist = CheckCookie();
  if (cookieExist) {
    axios({
      method:"post",
      url:"http://localhost:8080/show",
      data:bodyFormData,
      responseType:"json",
      withCredentials:true,
      headers:{ 'Content-Type' : 'multipart/form-data' }
    })
      .then(function (response) {
        var resp    = response.data;
        var r       = JSON.stringify(resp);
        var obj     = JSON.parse(r);
        var isEmpty = obj.IsEmpty;
        content     = obj.CredList;
        if (isEmpty) {
          table.empty = true ;
        } else if (content != null && !isEmpty) {
          table.empty = false;
          content.forEach(addToList)
          FillInTable()
          saveBtndisbl = false;
        }
      })
    category.value = "";
  }
}

function FillInTable() {
  var table = document.getElementById("table");
  var input = table.getElementsByTagName("input");
  if (input.length !== 0) {
    clearTable()
  }

  for  (var i=0;i<CNT.length;i++) {
    input.item(i).value = CNT[i];
  }
  CNT = []
}


function InsertTable(e) {
  e.preventDefault();
  GetTableData();
}

var UserLogged = CheckCookie();


function GrabTableData(){
  var tbody       = document.getElementById("myrows");
  var cnt         = tbody.getElementsByTagName("input");
  var category    = document.getElementById("ctg");
  var cookieExist = CheckCookie();

  payload.Category = filtered;
  for (var i=0;i<cnt.length;i++) {
    if ( cnt[i].value != "" ) {
      payload.Credential.push({"Username":cnt[i].value,"Password":cnt[i+1].value,"Category":cnt[i+2].value});
      i = i+2 ;
    } else if (cnt[i].value == "" && i != cnt.length) {
      continue;
    } else {
      table.updated = false;
      break;
    }
  }
  if (JSON.stringify(payload.Credential) !== JSON.stringify(content)) {
    PayloadContentCheck(payload.Credential)
    if (cookieExist && !fieldsRequire) {
      axios({
        method:"post",
        url:"http://localhost:8080/creds",
        data:JSON.stringify(payload),
        responseType:"json",
        withCredentials:true
      })
      .then(function (response) {
        var resp      = response.data;
        var r         = JSON.stringify(resp);
        var obj       = JSON.parse(r);
        table.updated = obj.Updated;
        visible       = true;
      })
    }
    payload.Credential = [];
  } else {
    payload.Credential = [];
  }
}

function PayloadContentCheck(content) {
  for (var i=0;i<content.length;i++) {
    if (content[i].Username === "" || content[i].Password === "" || content[i].Category === ""){
      fieldsRequire = true;
      visible       = true;
    } else {
      fieldsRequire = false;
    }
  }
}

function enableSave() {
  if (saveBtndisbl === true) {
    saveBtndisbl = false;
  }
}

function clicked(e) {
  var btn   = e.target;
  var id    = btn.id;
  var r     = /\d+/g;
  var ID;
  let nId;
  while ((ID=r.exec(id)) != null) {
    nId = ID[0]
  }
  var tr    = document.getElementById(nId);
  //var pwTocpy = document.getElementById(`pw${nId}`)
  var input = tr.getElementsByTagName("input");
  
  if (id === `edit${nId}`) {
    EditS(input);
  } else if (id === `show${nId}`){ 
    ShowS(input);
  } else if (id === `copy${nId}`) {
    Copy(input);
  }
}

function ShowS(el) {
  var pwf = el[1];
  if (pwf.type === "password") {
    pwf.type = "text";
  } else {
    pwf.type = "password"
  }
}

function EditS(el) {
  for (var i=0; i<el.length; i++) {
    if(el[i].disabled === true) {
      el[i].disabled = false ;
    } else {
      el[i].disabled = true;
    }
  }
}

function Copy(el) {
  var pwf = el[1];
  if (pwf.value !== "") {
    pwf.disabled = false ;
    pwf.type = "text";
    pwf.select();
    pwf.setSelectionRange(0,99999);
    document.execCommand("copy");
    pwf.disabled = true;
    pwf.type = "password";
    copied = true;
    visible= true;
  }
}




function AddS() {
  var table = document.getElementById("tbl");
  var row   = table.getElementsByTagName("tr");
  var newR  = row[row.length-1].cloneNode(true);
  var lines = row.length-1
  var rmBtn = document.getElementById("remove");
  newR.id   = lines
  var nInp  = newR.getElementsByTagName("input");
  for (var i=0;i<nInp.length;i++) {
    nInp[i].value = "";
    nInp[i].onkeyup = enableSave;
  }
  var nBtn  = newR.getElementsByTagName("button");
  for (var i=0;i<nBtn.length;i++) {
    if (i === 0) {
      nBtn[i].id      = `edit${lines}`
      nBtn[i].onclick = clicked;
    } else if (i === 1) {
      nBtn[i].id      = `show${lines}`
      nBtn[i].onclick = clicked;
    } else if (i === 2) {
      nBtn[i].id      = `copy${lines}`
      nBtn[i].onclick = clicked;
    }
  }
  table.appendChild(newR);
  rmBtn.disabled = false ;
}

function RemoveRow() {
  var table     = document.getElementById("tbl");
  var tr        = table.getElementsByTagName("tr");
  if ((tr.length-1) !== numObj*2) {
    var lastchild = table.lastChild;
    table.removeChild(lastchild);
  }
}


</script>

<main>
  { #if copied }
    <Alert color={succolor} isOpen={visible} toggle={() => (visible = false)}>
      Copied to clipboard!
    </Alert>
  { /if }
  { #if fieldsRequire }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
      Missed Some required fields in the Table
    </Alert>
  { /if }
  { #if table.updated }
    <Alert color={succolor} isOpen={visible} toggle={() => (visible = false)}>
      Saved!
    </Alert>
  { /if }
  { #if UserLogged }
    <div id="search">
      <form method="post" on:submit={InsertTable}>
        <input type="search" placeholder="Filter By Category" id="ctg">
        <Button  name="btn" class="sub" {btncolor}>Filter</Button>
      </form>
    </div>
    <div id="buttons">
    </div>
    <div id="table">
      <Table dark id="tbl">
        <thead>
          <tr>
            <th> Username </th>
            <th> Password </th>
            <th> Category </th>
            <th><Button {btncolor} id="add"  on:click={AddS}>+</Button></th>
            <th><Button {btncolor} id="remove" disabled=true on:click={RemoveRow}>-</Button></th>
            <th><Button {btncolor} id="Sv" disabled={saveBtndisbl} on:click={GrabTableData}>save</Button></th>
          </tr>
        </thead>
        <tbody id="myrows">
            { #each list as item, i }
              <tr id={i}>
                <td><input value="" type="text" on:keyup={enableSave} disabled=true></td>
                <td><input value="" id="pw{i}" type="password" on:keyup={enableSave} disabled=true></td>
                <td><input value="" type="text" on:keyup={enableSave} disabled=true></td>
                <td><Button {btncolor} id="edit{i}" on:click={clicked}> 🖊 </Button></td>
                <td><Button {btncolor} id="show{i}" on:click={clicked}>👁 </Button></td>
                <td><Button {btncolor} id="copy{i}" on:click={clicked}> 📋</Button></td>
              </tr>
            { /each }
        </tbody>
      </Table>
    </div>
  { :else }
    <Alert color={failcolor} isOpen={visible} toggle={() => (visible = false)}>
      Loggin To get Access To your Creds!
    </Alert>
  { /if }

</main>


<style>

  #buttons {
    padding-top:1.5%;
    float:right;
    padding-right:8%;
  }

  #search {
    padding-top:2%;
    padding-right:8%;
    padding-left:8%;
  }

  #table {
    padding-top: 2%;
    padding-left:6%;
    padding-right:8%;
  }
	.sub{
		/*background-color: #8cd7f4;*/
		color: black;
		padding: 16px 20px;
		margin: 8px 0;
		border: none;
		cursor: pointer;
		width: 50%;
		opacity: 0.9;
	}
	.sub:hover {
  	opacity:1;
	}

  input {
    font-weight: bold;
    color:black;
  }
</style>
