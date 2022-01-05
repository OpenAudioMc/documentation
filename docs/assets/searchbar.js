 function search_tags() {
    let input = document.getElementById('searchbar').value
    input=input.toLowerCase();
    let hide_title = document.getElementsByClassName('hide_on_search');
    let x = document.getElementsByClassName('search_tags');
    let count = 0;
    for (i = 0; i < x.length; i++) {
        if (!x[i].innerHTML.toLowerCase().includes(input)) {
        x[i].closest("div > a").parentElement.style.display="none";
        }
        else {
        x[i].closest("div > a").parentElement.style.display="initial";
        count+=1;
        }
    }
 if (input != ""){
    if (count > 0){
        document.getElementById("search_result_count").innerHTML = count + " results found."
    }
    else {
        document.getElementById("search_result_count").innerHTML = "no tags found, try again using another keyword."
    }
    for (i = 0; i < hide_title.length; i++) {
        hide_title[i].style.display="none";
        hide_title[i].parentElement.classList.remove("border-b-2","py-10");
        hide_title[i].parentElement.classList.add("py-1");
    }
 }
 else {
        document.getElementById("search_result_count").innerHTML = "";
        for (i = 0; i < hide_title.length; i++) {
            hide_title[i].style.display="flex";
            hide_title[i].parentElement.classList.add("py-1");
            hide_title[i].parentElement.classList.add("border-b-2","py-10");
        }
    }
}
//coded by someone awesome lol
