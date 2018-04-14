let globledata={};
let obj ={};
let obj_for_month={};
let obj_for_factory={};

function show_factory(){
    $.ajax({ //load進廠商清單
        url: "http://140.124.183.37:8067/GetIndex",
        type: "GET",
        success: function (msg) {
            var jsonObj = JSON.parse(msg);
            document.getElementById("manufac_list").innerHTML = "";
            console.log(jsonObj);

            for (var i = 0; i < jsonObj.length; i++) { //將取得的Json一筆一筆放入清單
                $("[id$=manufac_list]").append($("<option></option>").attr("value", jsonObj[i].Name).text(jsonObj[i].Name));
            }
        },
        error: function (xhr, ajaxOptions, thrownError) {
            alert(xhr.status);
        }
    });
} 

function show_factory_forUpload(){
    $.ajax({ //load進廠商清單
        url: "http://140.124.183.37:8067/GetIndex",
        type: "GET",
        success: function (msg) {
            var jsonObj = JSON.parse(msg);
            document.getElementById("manufac_list_upload").innerHTML = "";
            console.log(jsonObj);

            for (var i = 0; i < jsonObj.length; i++) { //將取得的Json一筆一筆放入清單
                $("[id$=manufac_list_upload]").append($("<option></option>").attr("value", jsonObj[i].ID).text(jsonObj[i].Name));
            }
        },
        error: function (xhr, ajaxOptions, thrownError) {
            alert(xhr.status);
        }
    });
} 

function doStatistics() { 

    var factorydate = document.selectForm.selectFactory.value;

    document.getElementById("tableTitle").innerHTML= factorydate + "年 資料庫統計報表";
    $.ajax({
        type: "GET",
        url: `http://140.124.183.37:8067/GetData?year=${factorydate}`
    }).done(function(msg) {
        $('#showData').show();
        $('#showChart').show();
        $('#forMonth').show();
        $('#forFactory').show();
        $('#choosePercentage').show(); 
        var factory_str = "";
        var json = msg;
        obj = JSON.parse(json);
        obj_for_month = JSON.parse(json);
        obj_for_factory = JSON.parse(json);

        document.getElementById('download1').onclick = () => JSONToCSVConvertor(json, factorydate, true);   
        document.getElementById("downloadChart2").onclick = () => downloadChart("showBar",factorydate);
        document.getElementById("downloadChart3").onclick = () => downloadChart("showLine",factorydate);
        document.getElementById("downloadChart4").onclick = () => downloadChart("showPie",factorydate);
        document.getElementById("downloadChart5").onclick = () => downloadChart("showBar_2",factorydate);
        document.getElementById("downloadChart6").onclick = () => downloadChart("showLine_f",factorydate);        

        document.getElementById('delete').onclick = () => Delete(factorydate);
        
        
        if (globledata.myBarChart != undefined){
            globledata.myBarChart.destroy();
            globledata.myBarChart=undefined;
        }           
        if (globledata.chart != undefined){
            globledata.chart.destroy();
            globledata.chart=undefined;
        } 
        if (globledata.myBarChart_2 != undefined){
            globledata.myBarChart_2.destroy();
            globledata.myBarChart_2=undefined;
        }    
        if (globledata.myBarChart_m != undefined){
            globledata.myBarChart_m.destroy();
            globledata.myBarChart_m=undefined;
        } 
        if (globledata.chart_m != undefined){
            globledata.chart_m.destroy();
            globledata.chart_m=undefined;
        } 
        if (globledata.myBarChart_f != undefined){
            globledata.myBarChart_f.destroy();
            globledata.myBarChart_f=undefined;
        } 
        if (globledata.lineChart_y != undefined){
            globledata.lineChart_y.destroy();
            globledata.lineChart_y=undefined;
        }   

        globledata.statistics = obj;  // add sort function
        orderStatistics("do not reorder");
        MonthlyGrowthRate(obj);
        toChart(obj);
        

        // type B
        // let headstr = '<thead><tr>' + keys.map(x=>`<th>${x}</th>`).join('') + '</tr></thead>';
        // let tablestr = headstr+'<tbody>';
        // for(let unit of obj){
        //     let bodystr = `<tr>${keys.map(x=>`<td>${unit[x]}</td>`).join('')}</tr>`;
        //     tablestr+=bodystr;
        // }
        // tablestr +='</tbody>';

        Chart.plugins.register({
            beforeDraw: function(chartInstance) {
              var ctx = chartInstance.chart.ctx;
              ctx.fillStyle = "white ";
              ctx.fillRect(0, 0, chartInstance.chart.width, chartInstance.chart.height);
            }
          });

    }).fail(err => {
        alert(err.responseText);
        window.location.reload();
    });  
     
    $.ajax({
        type: "GET",
        url: "http://140.124.183.37:8067/GetData/everyYearsTotalUsed"
    }).done(function(msg) {
        var factory_str = "";
        var json = msg;
        obj_year = JSON.parse(json);

        document.getElementById('download5').onclick = () => JSONToCSVConvertor(json, '全館資料庫每年使用增長情形', true);    
        
        var keys = [];
        var values = [];

        for(var i in obj_year){
            keys.push([(obj_year[i]['YEARS'])]);
            values.push([(obj_year[i]['Total'])]);
        }

        var tmpObj ={};
        var chartdataset =[];
        for(var j in obj_year){
            tmpObj ={};
            tmpObj.data = values[j];
            tmpObj.label = keys[j];
            tmpObj.backgroundColor = ('#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6) );
            chartdataset.push(tmpObj);
        }

        var ctx = document.getElementById("showBar_2").getContext("2d");
        globledata.myBarChart_2  = new Chart(ctx, {
                type: 'bar',
                data: {
                    // labels: keys,
                    datasets: chartdataset                  
                },
                options: {
                    title: {
                        display: true,
                        text: '全館資料庫每年使用增長情形'
                    },
                    scales: {
                        xAxes: [{ barPercentage: 0.5 }]
                    },
                    "animation": {
                        "duration": 1,
                        "onComplete": function() {
                          var chartInstance = this.chart,
                          ctx = chartInstance.ctx;
                  
                          ctx.fillStyle = "black";
                          ctx.textAlign = 'center';
                          ctx.textBaseline = 'bottom';
                  
                          this.data.datasets.forEach(function(dataset, i) {
                            var meta = chartInstance.controller.getDatasetMeta(i);
                            if(meta.hidden) return;
                            meta.data.forEach(function(line, index) {
                              var data = dataset.data[index];
                              ctx.fillText(data, line._model.x, line._model.y - 5);
                            });
                          });
                        }
                    }     
                }
        });

        document.getElementById('download6').onclick = () => JSONToCSVConvertor(obj_for_factory,'全館資料庫每月使用增長率', true);     

        var keys_y = [];
        var values_y = [];
    
        for(var i in obj_year){
            keys_y[i] = Object.keys(obj_year[i]);
            values_y[i] = Object.values(obj_year[i]);
        }

        for(var i in obj_year){
            keys_y[i].shift(); 
            keys_y[i].shift();            
            keys_y[i].pop(); 
            values_y[i].shift(); 
            values_y[i].shift();            
            values_y[i].pop(); 
        }  
    
        var tmpObj_y ={};
        var chartdataset_y =[];
        for(var j in obj_year){
            var color = ('#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6) );
            tmpObj_y = {};
            tmpObj_y.data = values_y[j];
            tmpObj_y.label = obj_year[j]['YEARS'];
            tmpObj_y.backgroundColor = color;        
            tmpObj_y.borderColor = color;            
            tmpObj_y.fill = false;
            chartdataset_y.push(tmpObj_y);
        }

        var chartdata_y = {
            labels: keys_y[0],
            datasets: chartdataset_y
        }
    
        var ctx_y = document.getElementById("showLine_f").getContext("2d");
        globledata.lineChart_y  = new Chart(ctx_y, {
            type: 'line',
            data: chartdata_y,
            options: {
                title: {
                    display: true,
                    text: '全館資料庫每月使用增長率'
                },
                "animation": {
                    "duration": 1,
                    "onComplete": function() {
                      var chartInstance = this.chart,
                      ctx = chartInstance.ctx;
                      console.log('test')
              
                    //   ctx.font = Chart.helpers.fontString(Chart.defaults.global.defaultFontSize, Chart.defaults.global.defaultFontStyle, Chart.defaults.global.defaultFontFamily);
                      ctx.textAlign = 'center';
                      ctx.textBaseline = 'bottom';
                      ctx.fillStyle = "black";
              
                      this.data.datasets.forEach(function(dataset, i) {
                        var meta = chartInstance.controller.getDatasetMeta(i);
                        if(meta.hidden) return;
                        meta.data.forEach(function(bar, index) {
                          var data = dataset.data[index];
                          ctx.fillText(data, bar._model.x, bar._model.y - 5);
                        });
                      });                      
                    }
                  },   
            }
        });
    }).fail(err => {
        alert(err.responseText);
        window.location.reload();
    });
} 

function orderStatistics(name){
    let compare = (x,y)=>x[name]>y[name]?-1:x[name]<y[name]?1:0; // arrow function 參數=>回傳值
    let obj = globledata.statistics;
    obj.sort(compare);

    let numPercentage = 0;
    let arrPercentage = [];

    for(let i in obj)
        numPercentage += obj[i]['Total'];
    
    for(let o of obj){
        o["%"]= (o['Total']*100/numPercentage).toFixed(2); 
    }

    let keys = Object.keys(obj[0]);
    let ignore = ['Total'];
    if(!document.getElementById('selectPer').checked)ignore.push('%');
    keys = keys.filter(x=>!ignore.includes(x));
    keys.push('Total')

    let name_map = {NAME:'單位名稱'}
    let rename = function(name){return name in name_map?name_map[name]:name}


    let sum = {};
    for(let key of keys){
        sum[key] = obj.map(o=>o[key]).reduce((a,b)=>a+b)
    }
    sum['NAME'] = '總計';
    sum['%'] = '100 %';

    var tablestr = "";
    tablestr +=
        `
        <thead>
            <tr> ${keys.map(x=>`<th class="${name==x?'sortorder':''}" onclick="orderStatistics('${x}');">${rename(x)}</th>`).join('')} </tr>
        </thead>
        <tbody>
            ${obj.map(unit=>`<tr>${keys.map(x=>`<td >${unit[x]}${x=='%'?' %':''}</td>`).join('')}</tr>`).join('')}
        <!-- sum -->
        <tr>
            ${keys.map(key=>`<td>${sum[key]}</td>`).join('')}
        </tr>
        </tbody>
        `
    document.getElementById("report").innerHTML = tablestr;

    

}
    
function Upload(){ //讀取csv檔案
    input = document.getElementById('fileinput');
    file = input.files[0];
    fr = new FileReader();
    fr.onload = receivedText;
    fr.readAsText(file);

    function receivedText() {
        submitForm(fr.result);
        // fr = new FileReader();
        // fr.onload = receivedBinary;
        // fr.readAsDataURL(file)
        // fr.readAsBinaryString(file);
    }

    // function receivedBinary() {
    //     // console.log(fr.result);
    //     // submitForm(fr.result);
    // }
}

function submitForm(data) { //上傳csv檔案
    console.log("submit event");

    var fileyear = document.fileinfo.selectFactory_year.selectedIndex;
    var fileyear_name = document.fileinfo.selectFactory_year.options[fileyear].value;
    
    var file_db = document.fileinfo.manufac_list_upload.selectedIndex;
    var file_db_name = document.fileinfo.manufac_list_upload.options[file_db].value;
    
    console.log(fileyear_name);
    console.log(file_db_name);

    $.ajax({
        url: "http://140.124.183.37:8067/Upload?year=" + fileyear_name + "&CompanyId=" + file_db_name, 
        method: "POST",
        data: data,
        headers:{
            "Content-Type": "text/plain",
        },
        processData: false,  // tell jQuery not to process the data
        contentType: "application/x-www-form-urlencoded", // tell jQuery not to set contentType
        error: function(xhr, ajaxOptions, thrownError) { //印出錯誤訊息
            alert(xhr.responseText);
            window.location.reload();
        },
        beforeSend:function(){
            $('#loading').show();
        }
    }).done(function( data ) { 
        $('#loading').hide();
        console.log( data );
        window.location.reload();
    });
    return false;
}

function MonthlyGrowthRate(obj){
    console.log(obj);
    var factorydate = document.selectForm.selectFactory.value;

    document.getElementById('download2').onclick = () => JSONToCSVConvertor(obj, factorydate+ '全館不同資料庫全文下載次數', true);                

    let keys = [];
    let values = [];

    let compare = (x,y)=>x['Total']>y['Total']?-1:x['Total']<y['Total']?1:0; // arrow function 參數=>回傳值
    obj.sort(compare);
    
    for(let i in obj){
        keys.push([(obj[i]['NAME'])]);
        values.push([(obj[i]['Total'])]);
    }  

    let tmpObj ={};
    let chartdataset =[];
    for(let j in obj){
        tmpObj ={};
        tmpObj.data = values[j];
        tmpObj.label = keys[j];
        tmpObj.backgroundColor = ('#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6));
        chartdataset.push(tmpObj);
    }        
    
    let chartdata = {
        datasets: chartdataset
    }

    let ctx = document.getElementById("showBar").getContext("2d");
    globledata.myBarChart  = new Chart(ctx, {
            type: 'bar',
            data: chartdata,
            options: {
                title: {
                    display: true,
                    text: factorydate + '全館不同資料庫全文下載次數'
                },
                scales: {
                    xAxes: [{ barPercentage: 0.4 }]
                },
                animation: {
                    onComplete: function () {
                      var chartInstance = this.chart;
                      var ctx = chartInstance.ctx;
                      ctx.textAlign = "center";
                      ctx.fillStyle = "black";                            
                      Chart.helpers.each(this.data.datasets.forEach(function (dataset, i) {
                        var meta = chartInstance.controller.getDatasetMeta(i);
                        if(meta.hidden) return;
                        Chart.helpers.each(meta.data.forEach(function (bar, index) {
                          ctx.save();
                          // Translate 0,0 to the point you want the text
                          ctx.translate(bar._model.x - 6 , bar._model.y - 27);
                  
                          // Rotate context by -90 degrees
                          ctx.rotate(-0.5 * Math.PI);
                  
                          // Draw text
                          ctx.fillText(dataset.data[index], 0, 0);
                          ctx.restore();
                        }),this)
                      }),this);
                    }
                  }
            }
    });
}

function show_chart(){

    if (globledata.lineChart != undefined){
        globledata.lineChart.destroy();
        globledata.lineChart=undefined;
    }   

    document.getElementById('download3').onclick = () => JSONToCSVConvertor(obj, factorydate + '每月使用增長率', true);   

    var factorydate = document.selectForm.selectFactory.value;    
    var section = document.selectSection.Section.value;    

    var new_obj = [];
    var one_part = Math.round( (Object.keys(obj).length/3 )); //取三等份
    var third = one_part * section ; 

    switch(section){
        case '1':
            new_obj = obj.slice(0,one_part);  
            console.log('1');
            break;
        case '2':
            new_obj = obj.slice(one_part,third);
            console.log('2');            
            break;
        case '3':
            new_obj = obj.slice(one_part*2,third);
            console.log('3');            
            break;        
        case '4':
            new_obj = obj;
            console.log('4');            
            break;  
    }

    var keys_m = [];
    var values_m = [];
    for(var i in new_obj){
        keys_m[i] = Object.keys(new_obj[i]);
        values_m[i] = Object.values(new_obj[i]);
    }

    for(var i in new_obj){
        keys_m[i].shift(); // name
        keys_m[i].pop(); // %
        keys_m[i].pop(); // total
        values_m[i].shift(); // name
        values_m[i].pop(); // %
        values_m[i].pop(); // total
    }
    
    var tmpObj_m ={};
    var chartdataset_m =[];
    for(var j in new_obj){
        tmpObj_m ={};
        tmpcolor = ('#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6) );
        tmpObj_m.data = values_m[j];
        tmpObj_m.label = new_obj[j]['NAME'];
        tmpObj_m.borderColor = tmpcolor;
        tmpObj_m.backgroundColor = tmpcolor;        
        tmpObj_m.fill = false;
        chartdataset_m.push(tmpObj_m);
    }

    var chartdata_m = {
        labels: keys_m[0],
        datasets: chartdataset_m
    }
    
    var ctx_m = document.getElementById("showLine").getContext("2d");
    globledata.lineChart = new Chart(ctx_m, {
        type: 'line',
        data: chartdata_m,
        options: {
            title: {
                display: true,
                text: factorydate +' 每月使用增長率'
            },
            "animation": {
                "duration": 1,
                "onComplete": function() {
                  var chartInstance = this.chart,
                  ctx = chartInstance.ctx;
          
                  ctx.font = Chart.helpers.fontString(Chart.defaults.global.defaultFontSize, Chart.defaults.global.defaultFontStyle, Chart.defaults.global.defaultFontFamily);
                  ctx.fillStyle = "black";
                  ctx.textAlign = 'center';
                  ctx.textBaseline = 'bottom';
          
                  this.data.datasets.forEach(function(dataset, i) {
                    var meta = chartInstance.controller.getDatasetMeta(i);
                    if(meta.hidden) return;
                    meta.data.forEach(function(line, index) {
                      var data = dataset.data[index];
                      ctx.fillText(data, line._model.x, line._model.y - 5);
                    });
                  });
                }
            },   
        }
    });    
}

function toChart(units){ //sort

    /// help functions
    let randomColor = ()=>'#'+ (0x1000000+Math.floor(Math.random()*(0xffffff+1))).toString(16).substr(1)

    //fetch necessary part
    let data = units.map(u=>({name:u['NAME'],value:u['Total'],color:randomColor()}))
    //order by total
    data.sort((a,b)=>a.value<b.value?1:a.value>b.value?-1:0)

    //data is ordered
    //prepare data for chart.js

    let value = data.map(x=>x.value)
    let total = value.reduce((x,y)=>x+y)
    let label = data.map(x=>`${x.name} ⬤ ${x.value} ⬤ ${(x.value*100/total).toFixed(2)}%`)
    let color = data.map(x=>x.color)

    let year = document.getElementById('selectFactory').value;
    let title = `${year} 年使用率`;
  
    document.getElementById('download4').onclick = () => JSONToCSVConvertor(units, year + '使用率', true); 
    
    //draw chart
    let ctx = document.getElementById('showPie').getContext('2d');
    globledata.chart = new Chart(ctx,{
        type: 'pie',
        data: {
            datasets: [{
                backgroundColor: color,
                data: value
            }],
            labels: label
        },
        options: {
            title: {display: true,text: title},
            tooltips:{callbacks:{label: tooltipItem => label[tooltipItem.index]}}
        }
    });
}

function JSONToCSVConvertor(JSONData, ReportTitle, ShowLabel) {
    //If JSONData is not an object then JSON.parse will parse the JSON string in an Object
    var arrData = typeof JSONData != 'object' ? JSON.parse(JSONData) : JSONData;
    
    var CSV = '';    
    //Set Report title in first row or line
    // CSV += ReportTitle + '\r\n\n';
    //This condition will generate the Label/Header
    if (ShowLabel) {
        var row = "";

        //This loop will extract the label from 1st index of on array
        for (var index in arrData[0]) {                    
            //Now convert each value to string and comma-seprated
            row += index + ',';
        }
        row = row.slice(0, -1);                
        //append Label row with line break
        CSV += row + '\r\n';
    }
    
    //1st loop is to extract each row
    for (var i = 0; i < arrData.length; i++) {
        var row = "";                
        //2nd loop will extract each column and convert it in string comma-seprated
        for (var index in arrData[i]) {
            row += '"' + arrData[i][index] + '",';
        }
        row.slice(0, row.length - 1);                
        //add a line break after each row
        CSV += row + '\r\n';
    }

    if (CSV == '') {        
        alert("Invalid data");
        return;
    }   
    
    //Generate a file name
    var fileName = "";
    //this will remove the blank-spaces from the title and replace it with an underscore
    fileName += ReportTitle.replace(/ /g,"_");   
    
    //Initialize file format you want csv or xls
    var uri = 'data:text/csv;charset=utf-8,\uFEFF' + encodeURI(CSV);
    // 原始版本，轉中文會亂碼，改成上面的版本
    // var uri = 'data:text/csv;charset=utf-8,' + escape(CSV);
    
    // Now the little tricky part.
    // you can use either>> window.open(uri);
    // but this will not work in some browsers
    // or you will not get the correct file extension    
    
    //this trick will generate a temp <a /> tag
    var link = document.createElement("a");    
    link.href = uri;
    
    //set the visibility hidden so it will not effect on your web-layout
    link.style = "visibility:hidden";
    link.download = fileName + ".csv";
    
    //this part will append the anchor tag and remove it after automatic click
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
}

function Delete(database){
    console.log(database);

    $.ajax({
        url: "http://140.124.183.37:8067/DataDelete?year=" + database,   //year ?????
        type: "GET",
        success: function (msg) {
            alert("Delete complete!")
            window.location.reload();
        },
        error: function (xhr, ajaxOptions, thrownError) {
            alert(xhr.status);
        }
    });
}

function downloadExample(){
    var filetype = document.fileinfo.upload.selectedIndex; //格式index
    switch(filetype){
        case 1:
            window.location.href = '../static/exampleData/format1.csv'; 
            break;
        case 2:
            window.location.href = '../static/exampleData/format2.csv'; 
            break;
        case 3:
            window.location.href = '../static/exampleData/format3.csv'; 
            break;
        case 4:
            window.location.href = '../static/exampleData/format4.csv'; 
            break;
        default:
            break;
    }
}

function downloadChart(chartid, name){
    var canvas = document.getElementById(chartid);

    var imgData = canvas.toDataURL("image/jpeg", 1.0);
    var pdf = new jsPDF("p", "mm", "a4");
    
    var width = pdf.internal.pageSize.width;    
    var height = pdf.internal.pageSize.height;
  
    pdf.addImage(imgData, 'JPEG', 0, 0, 207, 150);
    pdf.save(name +".pdf");

}

function doStatistics_m(){ 
    $('#showChart_m').show();

    var filetype_m = document.selectForm_m.selectMonth.selectedIndex;
    var filetype_month = document.selectForm_m.selectMonth.options[filetype_m].value;

    month = "";
    //轉換月份
    switch(filetype_month){
        case 'JAN':
            month = '1月份';
            break;
        case 'FEB':
            month = '2月份';
            break;
        case 'MAR':
            month = '3月份';
            break;
        case 'APR':
            month = '4月份';
            break;
        case 'MAY':
            month = '5月份';
            break;
        case 'JUN':
            month = '6月份';
            break;
        case 'JUL':
            month = '7月份';
            break;
        case 'AUG':
            month = '8月份';
            break;
        case 'SEP':
            month = '9月份';
            break;
        case 'OCT':
            month = '10月份';
            break;
        case 'NOV':
            month = '11月份';
            break;
        case 'DEC':
            month = '12月份';
            break;
    }

    document.getElementById('download7').onclick = () => JSONToCSVConvertor(obj_for_month, month + '全館不同資料庫全文下載次數', true);                
    document.getElementById("downloadChart7").onclick = () => downloadChart("showBar_m", month + '全館不同資料庫全文下載次數');        

    let keys_M = [];
    let values_M = [];

    let compare = (x,y)=>x[filetype_month] > y[filetype_month] ? -1 : x[filetype_month] < y[filetype_month] ? 1 : 0 ; // arrow function 參數=>回傳值
    obj_for_month.sort(compare);

    for(let i in obj_for_month){
        keys_M.push([(obj_for_month[i]['NAME'])]);
        values_M.push([(obj_for_month[i][filetype_month])]);
    }  

    let tmpObj_M ={};
    let chartdataset_M =[];
    for(let j in obj_for_month){
        tmpObj_M ={};
        tmpObj_M.data = values_M[j];
        tmpObj_M.label = keys_M[j];
        tmpObj_M.backgroundColor = ('#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6));
        chartdataset_M.push(tmpObj_M);
    }        

    let chartdata_M = {
        datasets: chartdataset_M
    }

    let ctx_M = document.getElementById("showBar_m").getContext("2d");
    globledata.myBarChart_m  = new Chart(ctx_M, {
            type: 'bar',
            data: chartdata_M,
            options: {
                title: {
                    display: true,
                    text: month + ' 全館不同資料庫全文下載次數'
                },
                scales: {
                    xAxes: [{ barPercentage: 0.4 }]
                },
                animation: {
                    onComplete: function () {
                      var chartInstance = this.chart;
                      var ctx = chartInstance.ctx;
                      ctx.textAlign = "center";
                      ctx.fillStyle = "black";                            
                      Chart.helpers.each(this.data.datasets.forEach(function (dataset, i) {
                        var meta = chartInstance.controller.getDatasetMeta(i);
                        if(meta.hidden) return;                        
                        Chart.helpers.each(meta.data.forEach(function (bar, index) {
                          ctx.save();
                          // Translate 0,0 to the point you want the text
                          ctx.translate(bar._model.x - 6 , bar._model.y - 27);
                  
                          // Rotate context by -90 degrees
                          ctx.rotate(-0.5 * Math.PI);
                  
                          // Draw text
                          ctx.fillText(dataset.data[index], 0, 0);
                          ctx.restore();
                        }),this)
                      }),this);
                    }
                  }
            }
    });

    document.getElementById('download8').onclick = () => JSONToCSVConvertor(obj_for_month, factorydate + ' 年' + month + '使用率' , true);  
    document.getElementById("downloadChart8").onclick = () => downloadChart("showPie_m", factorydate + ' 年' + month + '使用率');
    
    let randomColor = ()=>'#'+ (0x1000000+Math.floor(Math.random()*(0xffffff+1))).toString(16).substr(1)

    //fetch necessary part
    let data = obj_for_month.map(u=>({name:u['NAME'],value:u[filetype_month],color:randomColor()}))
    //order by total
    data.sort((a,b)=>a.value<b.value?1:a.value>b.value?-1:0)

    //data is ordered
    //prepare data for chart.js

    let value = data.map(x=>x.value)
    let total = value.reduce((x,y)=>x+y)
    let label = data.map(x=>`${x.name} ⬤ ${x.value} ⬤ ${(x.value*100/total).toFixed(2)}%`)
    let color = data.map(x=>x.color)

    let year = document.getElementById('selectFactory').value;
    let title = `${year} 年 ${filetype_month} 使用率`;

    var ctx_p = document.getElementById('showPie_m').getContext('2d');
    globledata.chart_m = new Chart(ctx_p,{
        type: 'pie',
        data: {
            datasets: [{
                backgroundColor: color,
                data: value
            }],
            labels: label,
            
        },
        options: {
            title: {
                display: true,
                text: title 
            },
            tooltips:{callbacks:{label: tooltipItem => label[tooltipItem.index]}}
        }
    });
}

function doStatistics_f(){
    $('#showChart_f').show();
    console.log(obj_for_factory)
    
    var factory = document.selectForm_f.manufac_list.selectedIndex;
    var factoryName = document.selectForm_f.manufac_list.options[factory].value;
    
    document.getElementById('download9').onclick = () => JSONToCSVConvertor(obj_for_factory, factoryName + '每月使用增長率', true);      
    document.getElementById("downloadChart9").onclick = () => downloadChart("showBar_f", factoryName + '每月使用增長率');
    
    var keys_f = [];
    var values_f = [];
    for(var i in obj_for_factory){
        keys_f = Object.keys(obj_for_factory[i]);
        if( factoryName == obj_for_factory[i]['NAME']){             
            values_f = Object.values(obj_for_factory[i]);    
        }
    }

    keys_f.shift();
    keys_f.pop();
    values_f.shift();
    values_f.pop();   

    var tmpbackgroundColor = [];
    for(var j in values_f){
        tmpbackgroundColor.push('#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6) );
    }

    var chartdata_f = {
        labels: keys_f,
        datasets: [
            {
              label: "使用次數",
              backgroundColor: tmpbackgroundColor,
              data: values_f
            }
          ]
    }
    
    var ctx_f = document.getElementById("showBar_f").getContext("2d");    
    globledata.myBarChart_f = new Chart(ctx_f, {
        type: 'bar',
        data: chartdata_f,
        options: {
            legend: { display: false },            
            title: {
                display: true,
                text: factoryName + ' 每月使用增長率'
            },
            "animation": {
                "duration": 1,
                "onComplete": function() {
                    var chartInstance = this.chart,
                        ctx = chartInstance.ctx;
            
                    ctx.font = Chart.helpers.fontString(Chart.defaults.global.defaultFontSize, Chart.defaults.global.defaultFontStyle, Chart.defaults.global.defaultFontFamily);
                    ctx.fillStyle = "black";                    
                    ctx.textAlign = 'center';
                    ctx.textBaseline = 'bottom';
            
                    this.data.datasets.forEach(function(dataset, i) {
                        var meta = chartInstance.controller.getDatasetMeta(i);
                        if(meta.hidden) return;                        
                        meta.data.forEach(function(bar, index) {
                        var data = dataset.data[index];
                        ctx.fillText(data, bar._model.x, bar._model.y - 5);
                        });
                    });
                }
            }                   
        }
    });


}