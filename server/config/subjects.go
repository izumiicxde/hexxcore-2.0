package config

import "hexxcore/types"

var PredefinedSchedule = []types.Schedule{
	{SubjectName: "ADA", DayOfWeek: "Wednesday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
	{SubjectName: "ADA", DayOfWeek: "Thursday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
	{SubjectName: "ADA", DayOfWeek: "Friday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
	{SubjectName: "ADA", DayOfWeek: "Saturday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
	{SubjectName: "IT", DayOfWeek: "Monday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
	{SubjectName: "IT", DayOfWeek: "Wednesday", StartTime: "11:45 AM", EndTime: "12:45 PM"},
	{SubjectName: "IT", DayOfWeek: "Friday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
	{SubjectName: "IT", DayOfWeek: "Saturday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
	{SubjectName: "SE", DayOfWeek: "Monday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
	{SubjectName: "SE", DayOfWeek: "Tuesday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
	{SubjectName: "SE", DayOfWeek: "Wednesday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
	{SubjectName: "SE", DayOfWeek: "Saturday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
	{SubjectName: "IC", DayOfWeek: "Wednesday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
	{SubjectName: "IC", DayOfWeek: "Thursday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
	{SubjectName: "LANG", DayOfWeek: "Monday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
	{SubjectName: "LANG", DayOfWeek: "Wednesday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
	{SubjectName: "LANG", DayOfWeek: "Thursday", StartTime: "11:45 AM", EndTime: "12:45 PM"},
	{SubjectName: "LANG", DayOfWeek: "Friday", StartTime: "11:45 AM", EndTime: "12:45 PM"},
	{SubjectName: "ENG", DayOfWeek: "Monday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
	{SubjectName: "ENG", DayOfWeek: "Tuesday", StartTime: "2:30 PM", EndTime: "3:30 PM"},
	{SubjectName: "ENG", DayOfWeek: "Wednesday", StartTime: "2:30 PM", EndTime: "3:30 PM"},
	{SubjectName: "ENG", DayOfWeek: "Friday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
	{SubjectName: "OE", DayOfWeek: "Tuesday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
	{SubjectName: "OE", DayOfWeek: "Thursday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
	{SubjectName: "ADA Lab", DayOfWeek: "Tuesday", StartTime: "9:30 AM", EndTime: "11:45 AM"},
	{SubjectName: "IT Lab", DayOfWeek: "Friday", StartTime: "1:30 PM", EndTime: "4:30 PM"},
}
