let globledata={};

function show_factory(){
    var factorydate = document.selectForm.selectFactory.value;
    $.ajax({ //load進廠商清單
        url: "http://140.124.183.37:8066/FileIndex?database="+ factorydate ,
        type: "GET",
        success: function (msg) {
            var jsonObj = JSON.parse(msg);
            document.getElementById("manufac_list").innerHTML = "";
            // console.log(jsonObj);
            //放入一筆空白的清單
            // $("[id$=manufac-list]").append($("<option></option>").attr("value", "").text(""));

            for (var i = 0; i < jsonObj.length; i++) {//將取得的Json一筆一筆放入清單
                $("[id$=manufac_list]").append($("<option></option>").attr("value", jsonObj[i].Name).text(jsonObj[i].Name));
            }
        },
        error: function (xhr, ajaxOptions, thrownError) {
            alert(xhr.status);
        }
    });
} 

function doStatistics() { 
    $('#showData').show();
    $('#showChart').show();
    $('#choosePercentage').show();    
    
    var factorydate = document.selectForm.selectFactory.value;
    var factory = document.selectForm.factoryList.selectedIndex;
    var factoryName = document.selectForm.factoryList.options[factory].value;
    document.getElementById("tableTitle").innerHTML= factorydate + "  _  " + factoryName + "  每月使用次數  ";
    $.ajax({
        type: "GET",
        url: `http://140.124.183.37:8066/GetData?database=${factorydate}&filename=${factoryName}`
    }).done(function(msg) {
        var factory_str = "";
        var json = msg;
        obj = JSON.parse(json);
        // console.log(obj);

        document.getElementById('download1').onclick = () => JSONToCSVConvertor(json, factoryName, true);   
        document.getElementById("downloadChart2").onclick = () => downloadChart("showBar");
        document.getElementById("downloadChart3").onclick = () => downloadChart("showLine");
        document.getElementById("downloadChart4").onclick = () => downloadChart("showBar_2");
        document.getElementById("downloadChart5").onclick = () => downloadChart("showPie");
        document.getElementById("downloadChart6").onclick = () => downloadChart("showPie_2");
        document.getElementById("downloadChart7").onclick = () => downloadChart("showPie_3");
        
        document.getElementById('delete').onclick = () => Delete(factorydate,factoryName);
        
        
        if (globledata.myBarChart != undefined){
            globledata.myBarChart.destroy();
            globledata.myBarChart=undefined;
        } 
        if (globledata.myBarChart_2 != undefined){
            globledata.myBarChart_2.destroy();
            globledata.myBarChart_2=undefined;
        }       
        if (globledata.lineChart != undefined){
            globledata.lineChart.destroy();
            globledata.lineChart=undefined;
        }       
        if (globledata.chart != undefined){
            globledata.chart.destroy();
            globledata.chart=undefined;
        } 
        if (globledata.chart2 != undefined){
            globledata.chart2.destroy();
            globledata.chart2=undefined;
        } 
        if (globledata.chart3 != undefined){
            globledata.chart3.destroy();
            globledata.chart3=undefined;
        } 

        toChart(obj); //繪製圓餅圖
        toChart_2(obj); //繪製學院
        
        globledata.statistics = obj;  // add sort function
        // console.log(obj);
        orderStatistics("do not reorder");
    
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
              ctx.fillStyle = "white";
              ctx.fillRect(0, 0, chartInstance.chart.width, chartInstance.chart.height);
            }
          });

    });

    var factory = document.selectForm.factoryList.selectedIndex;
    var factoryName = document.selectForm.factoryList.options[factory].value;
    $.ajax({
        type: "GET",
        url: `http://140.124.183.37:8066/GetData/Analysis?filename=${factoryName}`
    }).done(function(msg){
        let json = msg;
        let obj = JSON.parse(json);
        // console.log(obj);    

        document.getElementById('download2').onclick = () => JSONToCSVConvertor(json, factoryName+ '每年使用增長率', true);     
        document.getElementById('download3').onclick = () => JSONToCSVConvertor(json, factoryName+ '每月使用增長率', true);             

        let keys = [];
        let values = [];

        for(let i in obj){
            keys.push([(obj[i]['YEARS'])]);
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
            // labels: keys,
            datasets: chartdataset
        }

        let ctx = document.getElementById("showBar").getContext("2d");
        globledata.myBarChart  = new Chart(ctx, {
                type: 'bar',
                data: chartdata,
                options: {
                    title: {
                        display: true,
                        text: '每年使用增長率'
                    },
                    scales: {
                        xAxes: [{ barPercentage: 0.4 }]
                    },
                    "animation": {
                        "duration": 1,
                        "onComplete": function() {
                            var chartInstance = this.chart,
                                ctx = chartInstance.ctx;
                    
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
                    }    
                }
        });

        var keys_m = [];
        var values_m = [];
        for(var i in obj){
            keys_m[i] = Object.keys(obj[i]);
            values_m[i] = Object.values(obj[i]);
        }

        for(var i in obj){
            keys_m[i].shift(); // year
            keys_m[i].shift(); // name
            keys_m[i].pop(); // total
            values_m[i].shift(); // year
            values_m[i].shift(); // name
            values_m[i].pop(); // total
        }

        var tmpObj_m ={};
        var chartdataset_m =[];
        for(var j in obj){
            var color = ('#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6) );            
            tmpObj_m ={};
            tmpObj_m.data = values_m[j];
            tmpObj_m.label = obj[j]['YEARS'];
            tmpObj_m.backgroundColor = color;
            tmpObj_m.borderColor = color;        
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
                    text: '每月使用增長率'
                },
                "animation": {
                    "duration": 1,
                    "onComplete": function() {
                        var chartInstance = this.chart,
                            ctx = chartInstance.ctx;
                
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
                }                             
            }
        });
        
    });

    $.ajax({
        type: "GET",
        url: "http://140.124.183.37:8066/GetData/YearAnalysis",
        beforeSend:function(){
            $('#loading').show();
        }
    }).done(function(msg){
        $('#loading').hide();
        var json = msg;
        var obj = JSON.parse(json);
        // console.log(obj);  

        document.getElementById('download4').onclick = () => JSONToCSVConvertor(json, '全館資料庫使用增長情形', true);    

        var keys = [];
        var values = [];

        for(var i in obj){
            keys.push([(obj[i]['YEARS'])]);
            values.push([(obj[i]['Total'])]);
        }

        var tmpObj ={};
        var chartdataset =[];
        for(var j in obj){
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
                        text: '全館資料庫使用增長情形'
                    },
                    scales: {
                        xAxes: [{ barPercentage: 0.5 }]
                    },
                    "animation": {
                        "duration": 1,
                        "onComplete": function() {
                            var chartInstance = this.chart,
                                ctx = chartInstance.ctx;
                    
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
                    }      
                }
        });

    });  

    var factorydate = document.selectForm.selectFactory.value;
    $.ajax({
        type: "GET",
        url:`http://140.124.183.37:8066/GetData/DepartmentAnalysis?database=${factorydate}`
    }).done(function(msg){
        var json = msg;
        obj = JSON.parse(json);
        // console.log(obj);
        drawChartByYearInParticularSection(obj);
    });
} 

function orderStatistics(name){
    let compare = (x,y)=>x[name]>y[name]?-1:x[name]<y[name]?1:0; // arrow function 參數=>回傳值
    let obj = globledata.statistics;
    obj.sort(compare);

    let numPercentage = 0;
    let arrPercentage = [];

    if(!('%' in obj[0])){
        for(let i in obj)
            numPercentage += obj[i]['Total'];
    
        for(let o of obj){
            o['%']=Number((o['Total']*100/numPercentage).toFixed(2));
        }
    }


    // console.log(obj);
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
        fr = new FileReader();
        fr.onload = receivedBinary;
        fr.readAsBinaryString(file);
    }

    function receivedBinary() {
        submitForm(fr.result);
    }
}

function submitForm(data) { //上傳csv檔案
    console.log("submit event");
    // var fd = new FormData(document.getElementById("fileinfo"));
    var filetype = document.fileinfo.upload.selectedIndex;
    var filetype_name = document.fileinfo.upload.options[filetype].value;
    
    var filetype_db = document.fileinfo.selectFactory_ul.selectedIndex;
    var filetype_dbname = document.fileinfo.selectFactory_ul.options[filetype_db].value;
    
    var filetype_m = document.fileinfo.selectMonth.selectedIndex;
    var filetype_month = document.fileinfo.selectMonth.options[filetype_m].value;
    
    console.log(filetype_name);
    console.log(filetype_dbname);
    console.log(filetype_month);
    
    var filename = input.files[0];
    var filenamestr = filename.name;
    console.log(filenamestr);
    extIndex = filenamestr.lastIndexOf('.'); //去除檔名
    if (extIndex != -1) {
        filenamestr = filename.name.substr(0, extIndex);
    }
    console.log(filenamestr);
    $.ajax({
        url: "http://140.124.183.37:8066/"+ filetype_name + "?filename=" +filenamestr+ "&database=" + filetype_dbname + "&month=" + filetype_month,
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

function toChart(obj){ //sort

    document.getElementById('download5').onclick = () => JSONToCSVConvertor(obj, factorydate + ' 年 ' + factoryName + ' 使用率', true);  

    //Chart
    var tmpUnit = [];
    var tmpUnitTotal = [];
    var tmpColor = [];
    var colorcount = 0;
    var total = 0;
    var val = [];

    for(var id in obj)
        total += obj[id]['Total']
    for(var idx in obj)    
        val.push(Math.round(obj[idx]['Total'] / total * 100));    

    for(var unit_id in obj)
        tmpUnit.push(obj[unit_id]['NAME']);                                
    for(var unit_id in obj)
        tmpUnitTotal.push(obj[unit_id]['Total']);   

    for(var unit_id in obj)
        colorcount++;
    for(var i=0;i<=colorcount;i++){                
        tmpColor.push('#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6));         
    }

        // arrow function 參數=>回傳值
    var factorydate = document.selectForm.selectFactory.value;
    var factory = document.selectForm.factoryList.selectedIndex;
    var factoryName = document.selectForm.factoryList.options[factory].value;
    var ctx = document.getElementById('showPie').getContext('2d');
    globledata.chart = new Chart(ctx,{
        type: 'pie',
        data: {
            datasets: [{
                backgroundColor: tmpColor,
                data: tmpUnitTotal
            }],
            // These labels appear in the legend and in the tooltips when hovering different arcs
            labels: tmpUnit,
            
        },
        options: {
            title: {
                display: true,
                text: factorydate + ' 年 ' + factoryName + ' 使用率' 
            },
            tooltips: {
                callbacks: {
                    label: function(tooltipItem, data) {
                        var label = data.labels[tooltipItem.index];
                        var dataset = data.datasets[tooltipItem.datasetIndex];
                        var total = dataset.data.reduce(function(previousValue, currentValue, currentIndex, array) {
                            return previousValue + currentValue;
                        });
                        var currentValue = dataset.data[tooltipItem.index];
                        var precentage = Math.floor(((currentValue/total) * 100)+0.5);         
                        return label+ " : " + currentValue + " ( " + precentage + " % ) " ;
                    }
                }
            }
        }
    });
    
}

function toChart_2(obj){
    //Chart_2
    // console.log(obj);
    var tmpColor_2 = [];
    var colorcount_2 = 0;

    let lookup = {
        "機電學院":["機械系","車輛系","冷凍系","製科所"],
        "電資學院":["電機系","電子系","光電系","資工系"],
        "工程學院":["土木系","材資系","化工系","分子系"],
        "管理學院":["工工系","商管所經管系",],
        "設計學院":["建築系","工設系"],
        "人文與社會科學學院":["應用英文系"],
        "其他":["計網中心","共同科館","通識教育中心","技職教育所",
        "體育系","學務處 社團","總務處營繕組","行政大樓","國際事務處1~40",
        "圖書館","育成中心","語言中心","學院辦公室","EMBA辦公室","光電中心",
        "宿舍","中正館","自動化中心","教職員ADSL","無線網路","招生聯合會"],
        "校內未定義":["校內未定義"]
    }

    let inverse_lookup={};
    for(let 學院 in lookup){
        for(let 學系 of lookup[學院]){
            inverse_lookup[學系]=學院;
        }
    }

    let total={};

    for(let x of Object.keys(lookup)){
        total[x]=0;
    }

    let lose = new Set();
    for(let unit of obj){
        if(!inverse_lookup[unit['NAME']])lose.add(unit['NAME']); //not found
        else total[inverse_lookup[unit['NAME']]]+=unit['Total']; //found
    }

    keys = Object.keys(total);
    values = keys.map(x=>total[x]);
    // console.log(keys);

    let downloadData={};
    for(let i=0;i<keys.length;++i){
        downloadData[keys[i]]=values[i];
    }
    document.getElementById('download6').onclick = ()=> JSONToCSVConvertor([downloadData], factorydate + ' 年學院使用率', true)
        
    for(var unit_id in obj)
        colorcount_2++;
    for(var i=0;i<=colorcount_2;i++){                
        tmpColor_2.push('#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6));         
    }
    var factorydate = document.selectForm.selectFactory.value;
    var ctx = document.getElementById('showPie_2').getContext('2d');
    globledata.chart2 = new Chart(ctx,{
        type: 'pie',
        data: {
            datasets: [{
                backgroundColor: tmpColor_2,
                data: values
            }],
            // These labels appear in the legend and in the tooltips when hovering different arcs
            labels: keys
        },
        options: {
            title: {
                display: true,
                text: factorydate + ' 年學院使用率' 
            },
            tooltips: {
                callbacks: {
                    label: function(tooltipItem, data) {
                        var label = data.labels[tooltipItem.index];
                        var dataset = data.datasets[tooltipItem.datasetIndex];
                        var total = dataset.data.reduce(function(previousValue, currentValue, currentIndex, array) {
                            return previousValue + currentValue;
                        });
                        var currentValue = dataset.data[tooltipItem.index];
                        var precentage = Math.floor(((currentValue/total) * 100)+0.5);         
                        return label+ " : " + currentValue + " ( " + precentage + " % ) " ;
                    }
                }
            }
        }
    });
}

function drawChartByYearInParticularSection(obj){

    document.getElementById('download7').onclick = () => JSONToCSVConvertor(obj, factorydate + '年 年度單位使用率', true);    

    var tmpUnit = [];
    var tmpUnitTotal = [];
    var tmpColor = [];
    var colorcount = 0;

    for(var unit_id in obj)
        tmpUnit.push(obj[unit_id]['NAME']);                                
    for(var unit_id in obj)
        tmpUnitTotal.push(obj[unit_id]['Total']);   

    for(var unit_id in obj)
        colorcount++;
    for(var i=0;i<=colorcount;i++){                
        tmpColor.push('#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6));         
    }

        // arrow function 參數=>回傳值
    var factorydate = document.selectForm.selectFactory.value;
    var ctx = document.getElementById('showPie_3').getContext('2d');
    
    globledata.chart3 = new Chart(ctx,{
        type: 'pie',
        data: {
            datasets: [{
                backgroundColor: tmpColor,
                data: tmpUnitTotal
            }],
            // These labels appear in the legend and in the tooltips when hovering different arcs
            labels: tmpUnit
        },
        options: {
            title: {
                display: true,
                text: factorydate + '年 年度單位使用率'
            },
            tooltips: {
                callbacks: {
                    label: function(tooltipItem, data) {
                        var label = data.labels[tooltipItem.index];
                        var dataset = data.datasets[tooltipItem.datasetIndex];
                        var total = dataset.data.reduce(function(previousValue, currentValue, currentIndex, array) {
                            return previousValue + currentValue;
                        });
                        var currentValue = dataset.data[tooltipItem.index];
                        var precentage = Math.floor(((currentValue/total) * 100)+0.5);         
                        return label+ " : " + currentValue + " ( " + precentage + " % ) " ;
                    }
                }
            }
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

function Delete(database, filename){
    console.log(database);
    console.log(filename);

    $.ajax({
        url: "http://140.124.183.37:8066/DataDelete?database="+ database +"&filename=" + filename,
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

function downloadChart(chartid){
    var canvas = document.getElementById(chartid);

    var imgData = canvas.toDataURL("image/jpeg", 1.0);
    var pdf = new jsPDF("p", "mm", "a4");
    
    var width = pdf.internal.pageSize.width;    
    var height = pdf.internal.pageSize.height;
  
    pdf.addImage(imgData, 'JPEG', 0, 0, 207, 150);
    pdf.save(chartid +".pdf");

}
