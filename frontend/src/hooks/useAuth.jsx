import create from "zustand";

const useAuth = create((set) => ({
  setAuthToken: (authToken) => set(() => ({ authToken: saveToken(authToken) })),
}));

const saveToken = (token = "") => {
  localStorage.setItem("token", token);
};

export default useAuth;
