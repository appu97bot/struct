package main;
import "log"
func main(){
type person struct{
Name string;
height float32;
weight float32;
id string;
designation string;
}
a:=person{Name:"Appu",height:5.4,weight:62.3,id:"07A71A1297",designation:"housewife"};
l:=person{Name:"Samba",height:5.8,weight:67.5,id:"husband",designation:"IT employee"};
b:=person{Name:"Parvathi Surni",height:5.2,weight:70,id:"Mom",designation:"House head"};
c:=person{Name:"Venkateswarlu Surni",height:5.11,weight:77,id:"Dad",designation:"Railway employee"};
d:=person{Name:"Sandeep Surni",height:6.1,weight:80,id:"bro",designation:"IT Employee"};
e:=person{Name:"Sindhu",height:5.6,weight:65,id:"friendly sis",designation:"she is going to become like me"};
f:=person{Name:"Anu",height:5.5,weight:67,id:"sis",designation:"Mother"};
g:=person{Name:"Rani",height:5.2,weight:58,id:"Freind",designation:"IT employee"};
h:=person{Name:"Kalyani",height:5.1,weight:60,id:"friend",designation:"House wife"};
log.Println(person{"Aishwarya",5.6,65.7,"Aish","Actress"});
log.Println(person{"Abhishek",6.2,87.2,"abhi","Actor"});
log.Println("Person Details of a:",a);
log.Println("Person Details od l:",l);
log.Println("Person Details of b:",b);
log.Println("Person Details of c:",c);
log.Println("Person Details of d:",d);
log.Println("Person Details of e:",e);
log.Println("Person Details of f:",f);
log.Println("Person Details of g:",g);
log.Println("Person Details of h:",h);
a.Name="Aparna";
b.weight=70.7;
log.Println("Changed Details of a:",a);
log.Println("Changed Details of b:",b);
f=person{};
log.Println("Deleted details:",f);
f.Name="Kajol";
log.Println(f);
c.Name="Venkateswara Reddy";
log.Println(c);
}
