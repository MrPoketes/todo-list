import axios from 'axios';
import type { NextPage } from 'next';

interface ListPageProps {
	username: string;
}

const List: NextPage<ListPageProps> = ({ username }) => {
	console.log(username);
	return <div>S</div>;
};

export async function getServerSideProps(context: { query: { username: string } }) {
	console.log('context');
	console.log(context.query);
	return { props: { username: null } };
}

export default List;
