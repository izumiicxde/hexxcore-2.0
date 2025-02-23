"use client";
import Profile from "@/components/profile";
import { Checkbox } from "@/components/ui/checkbox";
import { DatePicker } from "@/components/ui/date-picker";
import { subjectStore } from "@/store/store";
import { ISubjectsResponse } from "@/types/classes";
import { useEffect, useState } from "react";
import { toast } from "sonner";

const Page = () => {
  const { subjects, setSubjects } = subjectStore();
  const [date, setDate] = useState<Date>(new Date());
  const [attendance, setAttendance] = useState<{
    [key: string]: boolean | null;
  }>({});

  const getSubjects = async () => {
    const url = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/attendance/subjects`;
    try {
      const response = await fetch(url, {
        credentials: "include",
        method: "GET",
      });

      const data: ISubjectsResponse = await response.json();
      toast(data.message);
      setSubjects(data.subjects);
    } catch (error) {
      toast("Unexpected error while fetching subjects");
    }
  };

  useEffect(() => {
    getSubjects();
  }, []);

  const handleCheckboxChange = (subject: string, status: boolean) => {
    setAttendance((prev) => ({
      ...prev,
      [subject]: prev[subject] === status ? null : status, // Toggle between attended, skipped, or unselected
    }));
  };

  const handleSubmit = async () => {
    if (!date) {
      toast("Please select a date");
      return;
    }

    const payload = {
      date: date.toISOString().split("T")[0], // Sends "2025-02-22"
      subjects: Object.entries(attendance)
        .filter(([_, status]) => status !== null) // Only send checked subjects
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
      <DatePicker date={date} setDate={setDate} />
      <div className="pt-5 w-full max-w-md">
        {subjects?.map((subject) => (
          <div
            className="flex justify-between items-center border p-3"
            key={subject}
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
        ))}
      </div>
      <button
        className="mt-4 px-6 py-2 bg-blue-500 text-white rounded"
        onClick={handleSubmit}
      >
        Submit Attendance
      </button>
    </div>
  );
};

export default Page;
