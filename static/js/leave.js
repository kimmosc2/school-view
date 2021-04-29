// 开始时间
let startHour = 8;
// 结束时间
let endHour = 20;
//请假事由
let reasonText = '';

let directionText = '';

let contacts = '';

let contactsTel = '';

function changTab1() {
    document.getElementById('tab1').setAttribute('class', 'tab active')
    document.getElementById('tab2').setAttribute('class', 'tab')
    document.getElementById('t_head1').setAttribute('class', 'tab-link  button active')
    document.getElementById('t_head2').setAttribute('class', 'tab-link   button')

}

function changeTab2() {
    document.getElementById('tab2').setAttribute('class', 'tab active')
    document.getElementById('tab1').setAttribute('class', 'tab')
    document.getElementById('t_head2').setAttribute('class', 'tab-link  button active')
    document.getElementById('t_head1').setAttribute('class', 'tab-link   button')
}

function openPop(){
    document.getElementById('popups').style.display = 'block'
}


function closePop(){
    document.getElementById('popups').style.display = 'none'
}