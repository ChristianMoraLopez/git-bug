import React from 'react';
import { Redirect, RouteComponentProps } from 'react-router-dom';

import CircularProgress from '@material-ui/core/CircularProgress';

import Bug from './Bug';
import { useGetBugQuery } from './BugQuery.generated';

type Props = RouteComponentProps<{
  id: string;
}>;

const BugQuery: React.FC<Props> = ({ match }: Props) => {
  const { loading, error, data } = useGetBugQuery({
    variables: { id: match.params.id },
  });
  if (loading) return <CircularProgress />;
  if (!data?.repository?.bug) return <Redirect to="/404bug" />;
  if (error) return <p>Error: {error}</p>;
  return <Bug bug={data.repository.bug} />;
};

export default BugQuery;
