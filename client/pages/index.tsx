import type { NextPage } from 'next';
import { Form } from '../components/Form/Form';
import { InputField } from '../components/InputField/InputField';
import { UnauthorizedLayout } from '../components/UnauthorizedLayout/UnauthorizedLayoutRoot';

const Home: NextPage = () => {
	return (
		<UnauthorizedLayout title="Login">
			<Form>
				<InputField name="username" label="Username" />
				<InputField name="password" type="password" label="Password" />
			</Form>
		</UnauthorizedLayout>
	);
};

export default Home;
