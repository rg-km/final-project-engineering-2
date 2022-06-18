import create from "zustand";
// import { devtools } from "zustand/middleware";

const useAuth = create((set) => ({
  authToken: "",
  setAuthToken: (authToken) => set(() => ({ authToken: authToken })),
}));

// const useAuth = create(devtools(auth));

export default useAuth;
