import create from "zustand";

const useAuth = create((set) => ({
  authToken: "",
  setAuthToken: (authToken) => set(() => ({ authToken: authToken })),
}));

export default useAuth;
