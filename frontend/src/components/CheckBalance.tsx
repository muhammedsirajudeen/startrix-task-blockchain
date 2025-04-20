import { CoffeeIcon } from "lucide-react";
import React, { Dispatch, SetStateAction, useState } from "react";
import {
    Dialog,
    DialogContent,
    DialogHeader,
    DialogTitle,
    DialogDescription,
    DialogClose,
} from "@/components/ui/dialog";

interface CheckBalanceProps {
    open: boolean;
    setOpen: Dispatch<SetStateAction<boolean>>;
}

const CheckBalance: React.FC<CheckBalanceProps> = ({ open, setOpen }) => {
    const [publicKey, setPublicKey] = useState("");
    const [balance, setBalance] = useState<number | null>(null);

    const handleCheckBalance = () => {
        setBalance(123.45);
    };

    return (
        <Dialog open={open} onOpenChange={setOpen}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>Check Coffee Coin Balance</DialogTitle>
                    <DialogDescription>
                        Enter your public key to check your balance.
                    </DialogDescription>
                </DialogHeader>
                <div className="mb-4">
                    <input
                        type="text"
                        value={publicKey}
                        onChange={(e) => setPublicKey(e.target.value)}
                        placeholder="Enter your public key"
                        className="w-full p-2 rounded-md border border-gray-300 focus:outline-none focus:ring-2 focus:ring-yellow-500"
                    />
                </div>
                <button
                    onClick={handleCheckBalance}
                    className="w-full bg-yellow-500 hover:bg-yellow-600 text-white font-bold py-2 px-4 rounded-md"
                >
                    Check Balance
                </button>
                {balance !== null && (
                    <div className="mt-4 flex items-center justify-center text-yellow-600 text-lg font-semibold">
                        <CoffeeIcon className="mr-2 text-yellow-500" />
                        <span>{balance} Coffee Coins</span>
                    </div>
                )}
                <DialogClose asChild>
                    <button className="mt-4 w-full bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-md">
                        Close
                    </button>
                </DialogClose>
            </DialogContent>
        </Dialog>
    );
};

export default CheckBalance;