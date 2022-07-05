import type { NextPage } from 'next';
import Link from 'next/link';
import { Button } from '../components/Button/Button';
import { Form } from '../components/Form/Form';
import { InputField } from '../components/InputField/InputField';
import { UnauthorizedLayout } from '../components/UnauthorizedLayout/UnauthorizedLayout';
import axios from 'axios';
import { useRouter } from 'next/router';

const Register: NextPage = () => {
	const router = useRouter();
	const handleRegister = async (values: Record<string, string>) => {
		const response = await axios.post(
			'http://localhost:9000/register',
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
			router.push(`/`);
		}
	};
	return (
		<>
			<UnauthorizedLayout.Root title="Register">
				<Form
					initialValues={{ username: '', password: '' }}
					onSubmit={handleRegister}
				>
					<InputField name="username" label="Username" required />
					<InputField
						name="password"
						type="password"
						label="Password"
						required
					/>
					<Button label="Register" />
				</Form>
				<UnauthorizedLayout.Footer>
					<Link href="/">
						<a>Log in</a>
					</Link>
				</UnauthorizedLayout.Footer>
			</UnauthorizedLayout.Root>
		</>
	);
};

export default Register;
