import React, { Dispatch, SetStateAction, useState } from "react";
import {
    Dialog,
    DialogContent,
    DialogHeader,
    DialogTitle,
    DialogFooter,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { toast } from "sonner";
import axiosInstance from "@/lib/axiosInstance";

interface AirdropDialogProps {
    open: boolean;
    setOpen: Dispatch<SetStateAction<boolean>>;
}

const AirdropDialog: React.FC<AirdropDialogProps> = ({ open, setOpen }) => {
    const [publicKey, setPublicKey] = useState("");

    const handleAirdrop = () => {
        try {
            if (publicKey.trim()) {
                console.log("Requesting airdrop for public key:", publicKey);
                axiosInstance.post('/airdrop',{address:publicKey})
                toast.success(<p className="text-white">Airdrop recieved successfully</p>,{style:{backgroundColor:"green"}})
            } else {
                toast.warning(<p>Public key is required</p>,{style:{backgroundColor:"orange"}});
            }            
        } catch (error) {
            console.log(error)
            toast.error(<p className="text-white" >Please try again</p>,{style:{backgroundColor:"red"}})
        }
    };

    return (
        <Dialog onOpenChange={setOpen} open={open}>
            <DialogContent className="bg-background text-foreground rounded-lg shadow-lg">
                <DialogHeader>
                    <DialogTitle className="text-xl font-semibold text-amber-500">
                        Request Airdrop
                    </DialogTitle>
                </DialogHeader>
                <div className="space-y-4">
                    <Input
                        placeholder="Enter your public key"
                        value={publicKey}
                        onChange={(e) => setPublicKey(e.target.value)}
                        className="border border-gray-300 focus:ring-amber-500 focus:border-amber-500 rounded-md"
                    />
                </div>
                <DialogFooter>
                    <Button
                        onClick={handleAirdrop}
                        className="bg-gradient-to-r from-amber-400 via-yellow-500 to-amber-600 text-white hover:opacity-90 rounded-full"
                    >
                        Submit
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    );
};

export default AirdropDialog;