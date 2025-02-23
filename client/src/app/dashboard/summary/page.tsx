"use client";
import { useEffect, useState } from "react";
import Profile from "@/components/profile";
import { toast } from "sonner";
import Link from "next/link";

interface SubjectSummary {
  subject_name: string;
  total: number;
  attended: number;
  skipped: number;
}

interface AttendanceSummary {
  total_classes: number;
  attended: number;
  skipped: number;
  allowed_skips: number;
  subjects: SubjectSummary[];
}

const SummaryPage = () => {
  const [summary, setSummary] = useState<AttendanceSummary | null>(null);
  const [loading, setLoading] = useState(true);

  const fetchSummary = async () => {
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_ENDPOINT}/attendance/summary`,
        {
          credentials: "include",
          method: "GET",
        }
      );

      if (!response.ok) {
        throw new Error("Failed to fetch data");
      }

      const data = await response.json();
      setSummary(data.summary); // Extracting "summary" from response
    } catch (error) {
      toast.error("Error fetching summary data");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchSummary();
  }, []);

  return (
    <div className="flex flex-col justify-center items-center w-full h-full p-5">
      <Profile />
      <Link href={"/dashboard/subjects"}>Mark attendance</Link>
      <h2 className="text-2xl font-bold mb-4">Attendance Summary</h2>

      {loading ? (
        <p>Loading summary...</p>
      ) : summary ? (
        <div className="mt-4 w-full max-w-lg p-6 border rounded-lg shadow-lg">
          <div className="grid grid-cols-2 gap-4 text-lg">
            <p>
              <strong>Total Classes:</strong> {summary.total_classes}
            </p>
            <p>
              <strong>Attended:</strong> {summary.attended}
            </p>
            <p>
              <strong>Skipped:</strong> {summary.skipped}
            </p>
            <p>
              <strong>Allowed Skips:</strong> {summary.allowed_skips}
            </p>
          </div>

          <h3 className="text-xl font-semibold mt-6 mb-3">Subject Breakdown</h3>
          <div className="space-y-3">
            {summary.subjects.map((subject) => (
              <div key={subject.subject_name} className="border p-4 rounded-lg">
                <p className="font-semibold text-lg">{subject.subject_name}</p>
                <p>Classes Taken: {subject.total}</p>
                <p>Attended: {subject.attended}</p>
                <p>Skipped: {subject.skipped}</p>
              </div>
            ))}
          </div>
        </div>
      ) : (
        <p className="mt-4 text-red-500">Failed to load summary.</p>
      )}
    </div>
  );
};

export default SummaryPage;
