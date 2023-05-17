import LogoutBtn from "../../components/LogoutBtn";
import { useGetUsersQuery } from "./usersApiSlice";
import { Link } from "react-router-dom";

const UsersList = () => {
    const {
        data: users,
        isLoading,
        isSuccess,
        isError,
        error,
    } = useGetUsersQuery("/users");
    let content;
    if (isLoading) {
        content = <p>"Loading..."</p>;
    } else if (isSuccess) {
        content = (
            <section className="users">
                <h1>Users List</h1>
                <ul>
                    {users.map((user: any, i: number) => {
                        return <li key={i}>{user.Username}</li>;
                    })}
                </ul>
                <Link to="/welcome">Back to Welcome</Link>
                <br />
                <Link to="/">Back to Public</Link>
                <LogoutBtn />
            </section>
        );
    } else if (isError) {
        content = (
            <section className="users">
                <h1>Users List</h1>
                <p>Unauthorized, You are not admin</p>
                <Link to="/welcome">Back to Welcome</Link>
                <br />
                <Link to="/">Back to Public</Link>
                <LogoutBtn />
            </section>
        );
    }

    return content || null;
};
export default UsersList;
