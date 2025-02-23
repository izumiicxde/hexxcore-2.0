"use client";
import Profile from "@/components/profile";
import { Checkbox } from "@/components/ui/checkbox";
import { DatePicker } from "@/components/ui/date-picker";
import { useEffect, useState } from "react";
import { toast } from "sonner";

const TodayPage = () => {
  const [subjects, setSubjects] = useState<string[]>([]);
  const [attendance, setAttendance] = useState<{
    [key: string]: boolean | null;
  }>({});
  const [date] = useState(new Date()); // Fixed to today’s date

  const getTodaySubjects = async () => {
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_ENDPOINT}/attendance/today`,
        { credentials: "include", method: "GET" }
      );

      const data = await response.json();
      console.log(data);
      setSubjects(data.subjects);
    } catch (error) {
      toast("Unexpected error while fetching today's subjects");
    }
  };

  useEffect(() => {
    getTodaySubjects();
  }, []);

  const handleCheckboxChange = (subject: string, status: boolean) => {
    setAttendance((prev) => ({
      ...prev,
      [subject]: prev[subject] === status ? null : status,
    }));
  };

  const handleSubmit = async () => {
    const payload = {
      date: date.toISOString().split("T")[0], // Sends "YYYY-MM-DD"
      subjects: Object.entries(attendance)
        .filter(([_, status]) => status !== null)
        .map(([name, status]) => ({ name, status })),
    };

    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_ENDPOINT}/attendance/mark`,
        {
          method: "POST",
          credentials: "include",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(payload),
        }
      );

      const data = await response.json();
      toast(data.message);

      if (response.ok) {
        setAttendance({});
      }
    } catch (error) {
      toast("Error submitting attendance");
    }
  };

  return (
    <div className="flex flex-col justify-center items-center w-full h-full p-5">
      <Profile />
      <h2 className="text-3xl font-bold p-5 ">Today's Attendance</h2>
      <DatePicker date={date} setDate={() => {}} disabled /> {/* Fixed date */}
      <div className="pt-5 w-full max-w-md">
        {subjects.length > 0 ? (
          subjects.map((subject) => (
            <div
              key={`${subject}`}
              className="flex justify-between items-center border p-3"
            >
              <span>{subject}</span>
              <div className="flex gap-3">
                <div className="flex items-center gap-1">
                  <Checkbox
                    checked={attendance[subject] === true}
                    onCheckedChange={() => handleCheckboxChange(subject, true)}
                  />
                  <span>Attended</span>
                </div>
                <div className="flex items-center gap-1">
                  <Checkbox
                    checked={attendance[subject] === false}
                    onCheckedChange={() => handleCheckboxChange(subject, false)}
                  />
                  <span>Skipped</span>
                </div>
              </div>
            </div>
          ))
        ) : (
          <p className="mt-4 text-gray-500">No classes today.</p>
        )}
      </div>
      {subjects.length > 0 && (
        <button
          className="mt-4 px-6 py-2 bg-blue-500 text-white rounded"
          onClick={handleSubmit}
        >
          Submit Attendance
        </button>
      )}
    </div>
  );
};

export default TodayPage;
