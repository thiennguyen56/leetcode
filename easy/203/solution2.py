class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def removeElements(self, head:Optional[ListNode], val:int) -> Optional[ListNode]:
        if head is None:
            return head
        
        temp_head = head
        prev_head = head

        while (temp_head):
            if temp_head.val == val:
                if temp_head == prev_head:
                    head = temp_head.next
                    prev_head = prev_head.next
                else:
                    prev_head = temp_head.next
            else:
                if temp_head != prev_head:
                    prev_head = prev_head.next
            temp_head = temp_head.next
        return head
