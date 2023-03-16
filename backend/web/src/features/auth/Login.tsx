import { useRef, useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import React from "react";
import { useDispatch } from "react-redux";
import { setCredentials } from "./authSlice";
import { useLoginMutation } from "../../app/api/authApiSlice";

const Login = () => {
    const userRef = useRef<HTMLInputElement>(null);
    const errRef = useRef<HTMLInputElement>(null);
    const [user, setUser] = useState("");
    const [pwd, setPwd] = useState("");
    const [errMsg, setErrMsg] = useState("");
    const navigate = useNavigate();

    const [login, { isLoading }] = useLoginMutation();
    const dispatch = useDispatch();

    useEffect(() => {
        userRef.current && userRef.current.focus();
    }, []);

    useEffect(() => {
        setErrMsg("");
    }, [user, pwd]);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        try {
            const userData = await login({ user, pwd }).unwrap();
            dispatch(setCredentials({ token: userData.accessToken, user }));
            setUser("");
            setPwd("");
            navigate("/welcome");
        } catch (err: any) {
            if (!err?.originalStatus) {
                // isLoading: true until timeout occurs
                setErrMsg("No Server Response");
            } else if (err.originalStatus === 400) {
                setErrMsg("Missing Username or Password");
            } else if (err.originalStatus === 401) {
                setErrMsg("Unauthorized");
            } else {
                setErrMsg("Login Failed");
            }
            errRef.current && errRef.current.focus();
        }
    };

    const handleUserInput = (e: React.ChangeEvent<HTMLInputElement>) =>
        setUser(e.target.value);

    const handlePwdInput = (e: React.ChangeEvent<HTMLInputElement>) =>
        setPwd(e.target.value);

    const content = isLoading ? (
        <h1>Loading...</h1>
    ) : (
        <section className="login">
            <p
                ref={errRef}
                className={errMsg ? "errmsg" : "offscreen"}
                aria-live="assertive"
            >
                {errMsg}
            </p>

            <h1>Employee Login</h1>

            <form onSubmit={handleSubmit}>
                <label htmlFor="username">Username:</label>
                <input
                    type="text"
                    id="username"
                    ref={userRef}
                    value={user}
                    onChange={handleUserInput}
                    autoComplete="off"
                    required
                />

                <label htmlFor="password">Password:</label>
                <input
                    type="password"
                    id="password"
                    onChange={handlePwdInput}
                    value={pwd}
                    required
                />
                <button>Sign In</button>
            </form>
        </section>
    );

    return content;
};
export default Login;
