package main;
import "log"
type person struct{
Name string;
height float32;
weight float32;
id string;
designation string;
}
func main(){
var a []person;
a=make([]person,10,10);
a[0]=person{Name:"Appu",height:5.4,weight:62.3,id:"07A71A1297",designation:"housewife"};
a[1]=person{Name:"Samba",height:5.8,weight:67.5,id:"husband",designation:"IT employee"};
a[2]=person{Name:"Parvathi Surni",height:5.2,weight:70,id:"Mom",designation:"House head"};
a[3]=person{Name:"Venkateswarlu Surni",height:5.11,weight:77,id:"Dad",designation:"Railway employee"};
a[4]=person{Name:"Sandeep Surni",height:6.1,weight:80,id:"bro",designation:"IT Employee"};
a[5]=person{Name:"Sindhu",height:5.6,weight:65,id:"friendly sis",designation:"she is going to become like me"};
a[6]=person{Name:"Anu",height:5.5,weight:67,id:"sis",designation:"Mother"};
a[7]=person{Name:"Rani",height:5.2,weight:58,id:"Freind",designation:"IT employee"};
a[8]=person{Name:"Kalyani",height:5.1,weight:60,id:"friend",designation:"House wife"};
log.Println(person{"Aishwarya",5.6,65.7,"Aish","Actress"});
log.Println(person{"Abhishek",6.2,87.2,"abhi","Actor"});
for i,k:=range a{
log.Println("Person Details of a:",i,k);
}
a[0].Name="Aparna";
a[3].weight=70.7;
log.Println("Changed Details of name:",a[0]);
log.Println("Changed Details of weight:",a[3]);
}
