import axios from 'axios';
import type { NextPage } from 'next';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { NextResponse } from 'next/server';
import { Button } from '../components/Button/Button';
import { Form } from '../components/Form/Form';
import { InputField } from '../components/InputField/InputField';
import { UnauthorizedLayout } from '../components/UnauthorizedLayout/UnauthorizedLayout';

const Home: NextPage = () => {
	const router = useRouter();
	const handleLogin = async (values: Record<string, string>) => {
		const response = await axios.post(
			'http://localhost:9000/login',
			{},
			{
				auth: {
					username: values.username,
					password: values.password
				}
			}
		);
		// Later show toast
		console.log(response.data);
		if (response.data.status === 200) {
			router.push(`/list/${response.data.data.username}`);
		}
	};
	return (
		<UnauthorizedLayout.Root title="Login">
			<Form initialValues={{ username: '', password: '' }} onSubmit={handleLogin}>
				<InputField name="username" label="Username" required />
				<InputField name="password" type="password" label="Password" required />
				<Button label="Login" />
			</Form>
			<UnauthorizedLayout.Footer>
				<Link href="/register">
					<a>Register</a>
				</Link>
			</UnauthorizedLayout.Footer>
		</UnauthorizedLayout.Root>
	);
};

export default Home;
