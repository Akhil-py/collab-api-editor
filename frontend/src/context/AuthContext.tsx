import { createContext, useContext } from 'react';

export const AuthContext = createContext({ user: null, login: () => {}, logout: () => {} });
export function useAuthContext() { return useContext(AuthContext); } 