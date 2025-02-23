import { IUser } from "@/types/user";
import { persist } from "zustand/middleware";
import { create } from "zustand";

interface IUserStore {
  user: IUser | null;
  setUser: (state: IUser) => void;
}

export const userStore = create<IUserStore>()(
  persist(
    (set) => ({
      user: null,
      setUser: (user) => set(() => ({ user })),
    }),
    { name: "user-storage" }
  )
);

interface ISubjectStore {
  subjects: string[] | null;
  setSubjects: (state: string[]) => void;
}

export const subjectStore = create<ISubjectStore>()(
  persist(
    (set) => ({
      subjects: null,
      setSubjects: (subjects) => set(() => ({ subjects })),
    }),
    { name: "subject-storage" }
  )
);
