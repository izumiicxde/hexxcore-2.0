"use client";

import { useEffect, useState } from "react";
import { toast } from "sonner";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
  ResponsiveContainer,
  PieChart,
  Pie,
  Cell,
} from "recharts";

const Dashboard = () => {
  const [summary, setSummary] = useState<any>(null);

  const fetchSummary = async () => {
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_ENDPOINT}/attendance/summary`,
        {
          credentials: "include",
          method: "GET",
        }
      );

      const data = await response.json();
      console.log("API Response:", data);
      setSummary(data.summary); // Correctly setting the summary object
    } catch (error) {
      toast("Error fetching summary data");
    }
  };

  useEffect(() => {
    fetchSummary();
  }, []);

  if (!summary) return <p>Loading...</p>;

  const barData = summary.subjects.map((subject: any) => ({
    name: subject.subject_name,
    Attended: subject.attended,
    Skipped: subject.skipped,
    Total: subject.total,
  }));

  const pieData = [
    { name: "Attended", value: summary.attended },
    { name: "Skipped", value: summary.skipped },
  ];

  const COLORS = ["#00C49F", "#FF4444"];

  return (
    <div className="flex flex-col items-center p-5 w-full">
      <h1 className="text-2xl font-bold mb-4">Attendance Dashboard</h1>

      {/* Total Summary */}
      <div className="w-full max-w-md border p-4 rounded-lg shadow-lg text-center mb-6">
        <h2 className="text-lg font-semibold">Overall Summary</h2>
        <p>Total Classes: {summary.total_classes}</p>
        <p>Attended: {summary.attended}</p>
        <p>Skipped: {summary.skipped}</p>
        <p>Allowed Skips: {summary.allowed_skips}</p>
      </div>

      {/* Bar Chart */}
      <div className="w-full max-w-2xl h-80">
        <h2 className="text-lg font-semibold text-center">
          Subject-Wise Attendance
        </h2>
        <ResponsiveContainer width="100%" height="100%">
          <BarChart data={barData}>
            <XAxis dataKey="name" />
            <YAxis />
            <Tooltip />
            <Legend />
            <Bar dataKey="Attended" fill="#00C49F" />
            <Bar dataKey="Skipped" fill="#FF4444" />
          </BarChart>
        </ResponsiveContainer>
      </div>

      {/* Pie Chart */}
      <div className="w-full max-w-md h-64 mt-6">
        <h2 className="text-lg font-semibold text-center">
          Overall Attendance
        </h2>
        <ResponsiveContainer width="100%" height="100%">
          <PieChart>
            <Pie
              data={pieData}
              dataKey="value"
              nameKey="name"
              cx="50%"
              cy="50%"
              outerRadius={100}
              label
            >
              {pieData.map((entry, index) => (
                <Cell key={`cell-${index}`} fill={COLORS[index]} />
              ))}
            </Pie>
            <Tooltip />
          </PieChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};

export default Dashboard;
