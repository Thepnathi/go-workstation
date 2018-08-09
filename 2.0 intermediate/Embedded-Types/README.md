# Notes

Difference between embedded type and traditional OOP subclassing (A class extends another and 
inherit all its methods and props). Parent and child goes both way. They both can access each other.

However, in Go, embedded type has no awareness of the tupe its embeded, so it it cannot access its fields or methods;
thus it only goes one way.