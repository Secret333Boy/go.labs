import { createContext } from "react";
import { Account } from "../types";

const defaultValue: Account = { Id: "" };

const AccountContext = createContext<Account>(defaultValue);

export default AccountContext;
