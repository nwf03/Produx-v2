import { useState } from 'react'
import { Button, Checkbox, Input, Modal, Text } from '@nextui-org/react'
import { useFollowProductMutation } from '../../state/reducers/api'
import { useRouter } from 'next/router'
export default function AccessPrivateProduct({
    show,
    setShow,
    productName,
    productIcon,
}: {
    show: boolean
    setShow: any
    productName: string
    productIcon: string | null
}) {
    const [showErrMsg, setShowErrMsg] = useState(false)
    const [token, setToken] = useState('')
    const closeHandler = () => {
        router.push('/')
        setShow(false)
    }
    const router = useRouter()
    const [follow, { error, isLoading }] = useFollowProductMutation()
    const handleFollow = async () => {
        await follow({ name: productName, follow: true, accessToken: token })
    }
    return (
        <div>
            <Modal preventClose blur aria-labelledby="modal-title" open={show}>
                {/*<Modal.Header className={"mt-2"}>*/}
                {/*  <Text id="modal-title" size={18}>*/}
                {/*    Follow*/}
                {/*    <Text b size={18} css={{ paddingLeft: 7 }}>*/}
                {/*      {productName}*/}
                {/*    </Text>*/}
                {/*  </Text>*/}
                {/*</Modal.Header>*/}
                <Modal.Body css={{ display: 'flex' }}>
                    {productIcon && (
                        <img
                            src={productIcon}
                            alt={productName}
                            className={'h-44 object-contain'}
                        />
                    )}
                    <Text
                        style={{ textAlign: 'center', marginBottom: 20 }}
                        size={18}
                    >
                        Follow
                        <Text b size={18} css={{ paddingLeft: 7 }}>
                            {productName}
                        </Text>
                    </Text>
                    <Input.Password
                        value={token}
                        onChange={(e) => setToken(e.target.value)}
                        clearable
                        bordered
                        fullWidth
                        color="primary"
                        size="lg"
                        placeholder="Access Token"
                    />
                </Modal.Body>
                <Modal.Footer>
                    <Button
                        auto
                        flat
                        color="error"
                        onClick={() => closeHandler()}
                    >
                        Go Back
                    </Button>
                    <Button auto onClick={() => handleFollow()}>
                        Follow
                    </Button>
                </Modal.Footer>
                {showErrMsg && (
                    <p className={'mb-4 text-red-500'}>
                        Please fill all the fields
                    </p>
                )}
                {error && (
                    <p className={'mb-4 text-red-500'}>{error.data.message}</p>
                )}
            </Modal>
        </div>
    )
}
