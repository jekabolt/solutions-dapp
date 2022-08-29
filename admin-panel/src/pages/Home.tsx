import { FC, useEffect } from 'react';
import type { AxiosError } from 'axios';
import { useQuery, useMutation } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-location';

import { 
  getNftRequests,
  getNftAllBurned,
  getNftAllBurnedError,
  getNftAllBurnedPending,
  uploadIpfsMetadata,
  burnNft,
  updateNftOffchainUrl,
  upsertNftMintRequest,
  uploadOffchainMetadata,
  updateBurnShippingStatus,
  deleteNftMintRequestById,
  deleteNftOffchainUrl,
  QUERIES,
  NFTMintRequestListArray,
} from 'api';
import { ROUTES } from 'constants/routes';
import { AUTH_LOCAL_STORAGE_KEY } from 'constants/values';

import styles from 'styles/HomePage.module.scss';

// +POST /api/auth/login
// -POST /api/nft/ipfs
// +POST /api/nft/burn
// +POST /api/nft
// +POST /api/nft/requests
// +POST /api/nft/offchain
// +POST /api/nft/shipping/status

// +DELETE /api/nft/requests/{id}
// +DELETE /api/nft/{id}

export const Home: FC = () => {
  const token = localStorage.getItem(AUTH_LOCAL_STORAGE_KEY) || '';
  const navigate = useNavigate();
  const { data, error, refetch } = useQuery<NFTMintRequestListArray, AxiosError>(
    [QUERIES.getNftRequests],
    () => getNftRequests(token),
    { enabled: false, retry: false },
  );

  const { refetch: refetch2 } = useQuery<any, AxiosError>(
    [QUERIES.getNftAllBurned],
    () => getNftAllBurned(token),
    { enabled: false, retry: false },
  );

  const { refetch: refetch3 } = useQuery<any, AxiosError>(
    [QUERIES.getNftAllBurnedError],
    () => getNftAllBurnedError(token),
    { enabled: false, retry: false },
  );

  const { refetch: refetch4 } = useQuery<any, AxiosError>(
    [QUERIES.getNftAllBurnedPending],
    () => getNftAllBurnedPending(token),
    { enabled: false, retry: false },
  );

  const { mutate: mutate1 } = useMutation(() => uploadIpfsMetadata(token));
  const { mutate: mutate2 } = useMutation(() => burnNft(
    token,
    {
      txid: 'wdwdwd',
      address: 'wdwd',
      mintSequenceNumber: 9991,
      shipping: {
        fullName: 'test',
        address: 'test',
        zipCode: 'test',
        city: 'test',
        country: 'test',
        email: 'test',
      }
    }
  ));
  const { mutate: mutate3 } = useMutation(() => uploadIpfsMetadata(token));
  const { mutate: mutate4 } = useMutation(() => upsertNftMintRequest(token,
    {
      sampleImages: [{ raw: 'base64' }],
      nftMintRequest: {
        id: 12345,
        ethAddress: 'ether address',
        TxHash: 'wefwefwe',
        mintSequenceNumber: 23,
        description: 'wdfwd'
      }
    }  
  ));
  const { mutate: mutate5 } = useMutation(() => uploadOffchainMetadata(token));
  const { mutate: mutate6 } = useMutation(() => updateBurnShippingStatus(token,
    { id: 'wdw', status: {trackNumber: 'wdwd', timeSent: 123456, error: 'wd1', success: true }}));
  const { mutate: mutate7 } = useMutation(() => deleteNftMintRequestById(token, { id: '12'}));
  const { mutate: mutate8 } = useMutation(() => deleteNftOffchainUrl(token, { id: '12idididi324' }));

  const test1 = () => {
    // /api/nft/requests
    refetch();
  };
  const test2 = () => {
    // /api/nft/burn
    refetch2();
  };
  const test3 = () => {
    // /api/nft/burn/error
    refetch3();
  };
  const test4 = () => {
    // /api/nft/burn/pending 
    refetch4();
  };


  const test5 = () => {
    // /api/nft/ipfs
    mutate1();
  };
  const test6 = () => {
    // /api/nft/burn
    mutate2();
  };
  const test7 = () => {
    // /api/nft
    mutate3();
  };
  const test8 = () => {
    // /api/nft/requests
    mutate4();
  };
  const test9 = () => {
    // /api/nft/offchain
    mutate5();
  };
  const test10 = () => {
    // /api/nft/shipping/status
    mutate6();
  };

  const test11 = () => {
    // /api/nft/shipping/status
    mutate7();
  };

  const test12 = () => {
    // /api/nft/shipping/status
    mutate8();
  };

  useEffect(() => {
    if (!token || error?.response?.status === 401) {
      navigate({ to: ROUTES.auth });
    } else {
      // if server stops infinity call occur
      // refetch();
    }
  }, [token, error?.response?.status]);

  return (
    <div className={styles.homePage}>
      home page
      <br />
      {data?.nftMintRequests?.map(() => (
        'single nft mint request'
      ))}
      <br />
      <br />
      <button onClick={() => localStorage.setItem(AUTH_LOCAL_STORAGE_KEY, '')}>clear localst</button>

      <br />
      <br />
      <br />
      <h3>get requests</h3>
      <br />
      <button onClick={test1}>/api/nft/requests</button>
      <br />
      <br />
      <button onClick={test2}>/api/nft/burn</button>
      <br />
      <br />
      <button onClick={test3}>/api/nft/burn/error</button>
      <br />
      <br />
      <button onClick={test4}>/api/nft/burn/pending</button>
      <br />
      <br />
      <br />
      <h3>post requests with test data</h3>
      <br />
      <br />
      <button onClick={test5}>/api/nft/ipfs</button>
      <br />
      <br />
      <button onClick={test6}>/api/nft/burn</button>
      <br />
      <br />
      <button onClick={test7}>/api/nft</button>
      <br />
      <br />
      <button onClick={test8}>/api/nft/requests</button>
      <br />
      <br />
      <button onClick={test9}>/api/nft/offchain</button>
      <br />
      <br />
      <button onClick={test10}>/api/nft/shipping/status</button>
      <br />
      <br />
      <br />
      <h3>delete requests with test data</h3>
      <br />
      <br />
      <button onClick={test11}>/api/nft/requests--id--</button>
      <br />
      <br />
      <button onClick={test12}>/api/nft/--id--</button>
    </div>
  );
};
